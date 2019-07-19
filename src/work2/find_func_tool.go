package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

//遍历指定目录中的代码，
// 并找出有“func () *dot.TypeLives”或“func () []*dot.TypeLives”定义的函数，
// 并生成调用这些函数的代码，
// 注：最好是能包含mod中的所有依赖
// gindot\kits.go\PreAddControlDot需要输入参数，这类函数是否需要排除？？
var path = flag.String("path", "", "find Absolute path,must be set")
var mod = flag.Bool("mod", true, "Traverse mod, default true")

//保存返回的组件通用信息
type packageInfo struct {
	absDir      string //该目录\包的绝对路径
	packageName string //正在扫描的包名，需要绝对路径
	astFiles    []*ast.File
	ImportDir   string //导入路径
	Alias       string //别名

	Funcs   []*FuncOfDot //返回组件普通信息的函数
	isExist bool         //是否存在返回dot.TypeLives的函数

	ConfigFuncNames []string //返回组件特有配置信息的函数
	IsExistConfig   bool     //是否存在返回config的函数
}
type FuncOfDot struct {
	FuncName string    //函数名
	IsSlice  bool      //返回值是否是数组
	location token.Pos //函数位置
}

func main() {

	var paths []string
	{ //命令行输入的顶级目录
		log.SetPrefix("find_func_tool: ") // log output Prefix
		flag.Usage = Usage                //使用自定义Usage替换默认提示
		flag.Parse()                      //解析命令行参数
		if len(*path) == 0 {
			flag.Usage()
			os.Exit(2)
		}
		//多个type参数时进行分割
		paths = strings.Split(*path, ",")
	}

	//处理mod文件得到依赖项目的路径
	for {
		if !isDir(paths) { //判断用户输入的目录能否找到
			os.Exit(2)
		}
		//都是能查到的目录
		var ifMod bool = *mod
		if ifMod { //用户是否选择使用mod
			paths = append(paths, getModFileData(paths)...)
		}
		if !isDir(paths) { //判断mod里面的依赖目录能否被找到
			os.Exit(2)
		}
		//得到绝对路径
		for ix := range paths {
			if !filepath.IsAbs(paths[ix]) {
				var err error
				paths[ix], err = filepath.Abs(paths[ix])
				if err != nil {
					log.Println(err, "构建绝对路径出错")
				}
			}
		}
		break
	}

	//保存目录以及子目录
	var allDirs []string
	{
		//获取所有的子目录
		for i := range paths {
			//遍历获取子目录
			dirs, err := getAllSonDirs(paths[i])
			if err != nil {
				log.Fatal("getAllSonDirs:", err)
			}
			allDirs = append(allDirs, dirs...)
		}
	}

	//处理重复的目录，原因是多个mod中可能会重复依赖多个项目
	// 或者a\b\c和a\b这类目录也要处理
	{
		allDirs = RemoveRepByMap(allDirs)
	}

	// 对于每个子目录通过package.config获取包名，以及当前位置的所有ast.file
	var allInfo []*packageInfo
	{
		//加载路径获取配置信息
		cfg := &packages.Config{
			Mode: packages.LoadSyntax, //不包含依赖,尝试下面这个
			//Mode: packages.LoadAllSyntax,
		}
		pkginfos, err := packages.Load(cfg, allDirs...)
		if err != nil {
			log.Fatal("package.Load:", err)
		}
		for ix := range allDirs {
			pkg := pkginfos[ix]
			astFiles := make([]*ast.File, len(pkg.Syntax))
			for i, file := range pkg.Syntax {
				astFiles[i] = file
			}
			p := packageInfo{
				absDir:      allDirs[ix],
				packageName: pkg.Name,
				astFiles:    astFiles,
				ImportDir:   pkg.PkgPath,
			}
			allInfo = append(allInfo, &p)
		}
	}

	//查找[]*packageInfo的每一个对象，根据ast,node判断
	{
		for _, p := range allInfo {
			p.findFuncNodeOnAst(false) //查找*dot.TypeLives和[]*dot.TypeLives
			p.findFuncNodeOnAst(true)  //查找*dot.ConfigTypeLives
		}
	}

	//将满足普通条件的包筛选出来
	var exitFuncInfos []*packageInfo
	{
		//为isExist字段赋值
		{
			for _, p := range allInfo {
				if len(p.Funcs) == 0 {
					//这个目录下没有满足条件的函数
					p.isExist = false
				} else {
					p.isExist = true
				}
			}
		}
		//赋值
		{
			for _, p := range allInfo {
				if p.isExist {
					exitFuncInfos = append(exitFuncInfos, p)
				}
			}
		}

		//检测有没有结果
		{
			if len(exitFuncInfos) == 0 {
				log.Fatal("没有找到符合条件的函数")
			}
		}
	}

	//解决同名包问题-别名
	{
		//存储每个包出现的次数
		map1 := make(map[string]int)
		//赋值
		{
			for _, p := range exitFuncInfos {
				//该包是否已经放入
				if _, ok := map1[p.packageName]; ok {
					//已有
					map1[p.packageName]++
				} else {
					//没有
					map1[p.packageName] = 1
				}
			}
		}
		//利用map1构建别名，只出现一次别名默认为包名
		{
			for _, p := range exitFuncInfos {
				if v, ok := map1[p.packageName]; ok {
					if v == 1 {
						p.Alias = p.packageName
					} else {
						p.Alias = p.packageName + "_" + strconv.Itoa(v)
						map1[p.packageName]--
					}
				}
			}
		}
	}

	//为isExistConfig字段赋值
	{
		for _, p := range exitFuncInfos {
			if len(p.ConfigFuncNames) == 0 {
				//这个目录下没有满足条件的函数
				p.IsExistConfig = false
			} else {
				p.IsExistConfig = true
			}

		}
	}

	//生成代码文件
	{
		buildCodeFromTemplate(exitFuncInfos)

	}

	//错误处理还没做
	//调用执行callMethod生成序列化的json文件
	{
		cmd := exec.Command(getGOROOTBin(), "run", "callMethod.go")
		//cmd := exec.Command("test.bat")
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error %v executing command!", err)
			os.Exit(1)
		}
	}
	//反序列化得到callMethod的结果
	//var reult = make([]*dot.TypeLives, 0)
	/*	var f interface{}
		{
			buf, err := ioutil.ReadFile("resultConfig.json")
			if err != nil {
				log.Fatal("读取json出错")
			}
			err = json.Unmarshal(buf, &f)
			if err != nil {
				log.Fatal("反序列化出错")
			}
			fmt.Println(f)

		}*/

}

//自定义使用提示
func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of find_func_tool:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tfind_func_tool [flags] -path=dir1,dir2 \n\t# Must be set\n")
	_, _ = fmt.Fprintf(os.Stderr, "For more information, see:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\thttps://github.com/scryinfo/demo/blob/master/generate/rpccall/findFuncTool/readme.md\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults() //print flags
}

func isDir(paths []string) bool {
	for i := range paths {
		//获取用户输入的目录
		var dir = paths[i]
		//判断用户输入是否正确
		if !isDirectory(dir) {
			log.Println("The directory of %v could not be found", dir)
			return false
		}
	}
	return true
}
func isDirectory(name string) bool {
	info, err := os.Stat(name) //return fileinfo
	if err != nil {
		log.Fatal("isDirectory:", err)
	}
	return info.IsDir() //true or false
}

//传入用户输入的目录，返回mod文件里面包含的所有内容
func getModFileData(paths []string) []string {

	var modFiles []string
	var modFileDirData []string
	//获取所有mod文件
	{
		for i := range paths {
			//获取用户输入的每一个目录包含的所有mod文件
			mod, err := getAllModFiles(paths[i])
			if err != nil {
				log.Fatal("getALlModFiles:", err)
			}
			modFiles = append(modFiles, mod...)
		}
	}
	//读出mod文件的内容
	{
		for i := range modFiles {
			modFileDirData = append(modFileDirData, getDirsFromMod(modFiles[i])...)
		}
	}
	return modFileDirData
}

// 读取指定目录下所有mod文件
func getAllModFiles(path string) (files []string, err error) {
	var suf = ".mod"
	var dirs []string
	{
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}
		PthSep := string(os.PathSeparator)
		for i := range dir {
			fi := dir[i]
			if fi.IsDir() { //目录下的目录文件
				dirs = append(dirs, filepath.Join(path, fi.Name())) //JoinPath, filepath.Join
			} else {
				// 过滤指定格式
				ok := strings.HasSuffix(fi.Name(), suf)
				if ok {
					files = append(files, path+PthSep+fi.Name())
				}
			}
		}
	}
	// 获取子目录下的mod文件
	for j := range dirs {
		xPath := dirs[j]
		xFiles, _ := getAllModFiles(xPath)
		files = append(files, xFiles...)
	}
	return files, nil
}

//每次读一个mod文件，获取里面的目录集合
func getDirsFromMod(path string) []string {
	datas, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("read mod file error，please check path or go mod download")
	}
	str := string(datas)
	start := "require ("
	end := ")"
	//获取内容
	{
		n := strings.Index(str, start)
		if n == -1 {
			n = 0
		}
		str = string([]byte(str)[n+len(start)+2:])
		m := strings.Index(str, end)
		if m == -1 {
			m = len(str)
		}
		str = string([]byte(str)[:m-1])
	}
	//获取内容中的目录
	var dirs []string
	{
		//按行读取
		var dataline = strings.Split(str, "\n")
		for i := range dataline {
			//按空格分割
			data := strings.Fields(dataline[i])
			//将导入路径的/替换为当前操作系统的路径分隔符
			data[0] = filepath.FromSlash(data[0])
			dir := getGOPATHsrc() + data[0]
			dirs = append(dirs, dir)
		}
	}
	return dirs
}
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for i := range slc {
		l := len(tempMap)
		tempMap[slc[i]] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, slc[i])
		}
	}
	return result
}

//获取指定目录下的所有子目录
func getAllSonDirs(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}
			return nil
		})
	return dir_list, dir_err
}

//查找满足条件的函数节点
func (p *packageInfo) findFuncNodeOnAst(isConfig bool) {
	///var FuncNames []string
	for _, astFile := range p.astFiles {
		if astFile == nil {
			break
		}
		ast.Inspect(astFile, func(node ast.Node) bool {
			for {
				//must be a func
				funcNode, ok := node.(*ast.FuncDecl)
				if !ok {
					break //不是函数
				}
				if funcNode.Recv != nil {
					break //排除有接收者的函数
				}
				if !(funcNode.Type.Params.List == nil) {
					break //排除需要参数的函数
				}
				result := funcNode.Type.Results
				if result == nil {
					break //排除没有返回值的函数
				}
				resultList := result.List
				if len(resultList) != 1 {
					break //排除返回值不值一个的函数
				}
				if isConfig {
					b1, _ := returnValueJudgmentOfConfig(resultList[0])
					if !b1 {
						break //排除返回值类型不匹配的函数
					}
					//保存函数信息
					p.ConfigFuncNames = append(p.ConfigFuncNames, funcNode.Name.Name)
				} else {
					b1, b2 := returnValueJudgment(resultList[0])
					if !b1 {
						break //排除返回值类型不匹配的函数
					}
					//保存函数信息
					funcInfo := &FuncOfDot{
						FuncName: funcNode.Name.Name, //函数名
						IsSlice:  b2,
						location: funcNode.Type.Func,
					}
					p.Funcs = append(p.Funcs, funcInfo)
				}
				return true
			}
			return true
		})
	}
}

//查找通用组件信息
//第一个代表函数的返回值是否符合条件
//第二个代表返回值是否是数组
func returnValueJudgment(ret *ast.Field) (bool, bool) {
	retype, ok := (ret.Type).(*ast.StarExpr) //找到*
	if ok {                                  //是一个指针
		x, ok1 := (retype.X).(*ast.SelectorExpr) //有选择器的表达式  a.b
		if !ok1 {
			return false, false //类似*xxx
		}
		xx := x.X.(*ast.Ident)
		xsel := x.Sel.Name
		if xx.Name == "dot" {
			if xsel == "TypeLives" {
				return true, false //返回值是*dot.Typelives
			}
		}
		return false, false //指针指向的结构错误
	}

	retype2, ok := (ret.Type).(*ast.ArrayType)
	if ok { //是一个切片
		elt, ok := (retype2.Elt).(*ast.StarExpr)
		if ok { //切片存放的指针数据
			x, ok1 := (elt.X).(*ast.SelectorExpr) //有选择器的表达式  a.b
			if !ok1 {
				return false, false //类似*xxx
			}
			xx := x.X.(*ast.Ident)
			xsel := x.Sel.Name
			if xx.Name == "dot" {
				if xsel == "TypeLives" {
					return true, true //返回值是[]*dot.Typelives
				}
			}
			return false, false //指针指向的结构错误
		}
		return false, false //切片存放的数据不是指针
	}
	return false, false //返回值类型错误
}

//查找特有组件信息
//第一个代表函数的返回值是否符合条件
//第二个代表返回值是否是数组
func returnValueJudgmentOfConfig(ret *ast.Field) (bool, bool) {
	retype, ok := (ret.Type).(*ast.StarExpr) //找到*
	if ok {                                  //是一个指针
		x, ok1 := (retype.X).(*ast.SelectorExpr) //有选择器的表达式  a.b
		if !ok1 {
			return false, false //类似*xxx
		}
		xx := x.X.(*ast.Ident)
		xsel := x.Sel.Name
		if xx.Name == "dot" {
			if xsel == "ConfigTypeLives" {
				return true, false //返回值是*dot.ConfigTypelives
			}
		}
		return false, false //指针指向的结构错误
	}
	return false, false //返回值类型错误
}

//
func getGOPATHsrc() string {
	gopath := os.Getenv("GOPATH")
	switch runtime.GOOS {
	case "windows":
		gopath = gopath + "\\src\\"
	case "linux":
		gopath = gopath + "/src/"
	default:
		log.Fatal("无法识别的操作系统")
	}
	return gopath
}
func getGOROOTBin() string {
	gopath := os.Getenv("GOROOT")
	switch runtime.GOOS {
	case "windows":
		gopath = gopath + "\\bin\\go.exe"
	case "linux":
		gopath = gopath + "/bin/"
	default:
		log.Fatal("无法识别的操作系统")
	}
	return gopath
}

//模板生成
func buildCodeFromTemplate(e []*packageInfo) {
	buf := bytes.Buffer{}
	//使用模板
	//fmt.Println(e)
	t, err := template.ParseFiles("file1.tmpl")
	if err != nil {
		log.Println("parseFileErr:", err)
	}
	err = t.Execute(&buf, e)
	if err != nil {
		log.Println("executing template:", err)
	}
	//file
	baseName := "callMethod.go"
	baseName = filepath.Join(".", baseName)
	err = ioutil.WriteFile(baseName, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

/*//字符串拼接生成
func buildCodeUseString(exitFuncInfo []*packageInfo) {
	buf := bytes.Buffer{}
	{
		//头部
		_, _ = fmt.Fprintf(&buf, "// Scry Info.  All rights reserved.\n")
		_, _ = fmt.Fprintf(&buf, "// license that can be found in the license file.\n")
		_, _ = fmt.Fprintf(&buf, "package main\n\n")
		_, _ = fmt.Fprintf(&buf, "import (\n")
	}
	{
		//import
		for _, e := range exitFuncInfo {
			_, _ = fmt.Fprintf(&buf, "\t%s \"%s\"\n", e.Alias, e.ImportDir)
		}
		_, _ = fmt.Fprintf(&buf, "\t)\n\n")
	}
	{
		//func
		_, _ = fmt.Fprintf(&buf, "func main() {\n")
		for _, e := range exitFuncInfo {
			for i := range e.FuncNames {
				_, _ = fmt.Fprintf(&buf, "\t%s.%s()\n", e.Alias, e.FuncNames[i])
			}
		}
		_, _ = fmt.Fprintf(&buf, "}")
	}
	//file
	baseName := "v1.go"
	baseName = filepath.Join(".", baseName)
	err := ioutil.WriteFile(baseName, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
*/
//解析子目录，为对象添加packagename以及ast.file信息
//nil
func (p *packageInfo) parsePackage() {
	//加载路径获取配置信息
	cfg := &packages.Config{
		Mode: packages.LoadSyntax, //不包含依赖,尝试下面这个
		//Mode: packages.LoadAllSyntax,
	}
	pkginfos, err := packages.Load(cfg, p.absDir)
	if err != nil {
		log.Fatal("package.Load:", err)
	}
	pkg := pkginfos[0]
	p.packageName = pkg.Name
	p.astFiles = make([]*ast.File, len(pkg.Syntax))
	for i, file := range pkg.Syntax {
		p.astFiles[i] = file
	}
}
