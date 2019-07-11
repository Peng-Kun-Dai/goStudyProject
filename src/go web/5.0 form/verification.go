package main

//验证表单的输入
//服务器端的验证
import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

func main() {
	//必填字段
	if len(r.Form["username"][0]) == 0 {
		//为空的处理
		r.From.Get()
	}

	//数字
	//第一种正则匹配
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		return false
	}
	//第二种
	getint, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		//数字转化出错了，可能不是数字
	}
	if getint > 100 {
		//太大了
	}

	//验证中文
	//Unicode包
	//func Is(rangeTab *RangeTable,r rune) bool
	//正则
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		return false
	}

	//验证英文
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}

	//手机号码
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		return false
	}

	//下拉菜单
	/*我们的select可能是这样的一些元素
	<select name="fruit">
	<option value="apple">apple</option>
	<option value="pear">pear</option>
	<option value="banane">banane</option>
	</select>*/
	//避免伪造数据
	slice := []string{"apple", "pear", "banane"}
	for _, v := range slice {
		if v == r.Form.Get("fruit") {
			return true
		}
	}
	//return false

	//单选按钮
	/*<input type="radio" name="gender" value="1">男
	<input type="radio" name="gender" value="2">女*/
	slice1 := []int{1, 2}

	for _, v := range slice1 {
		if v == r.Form.Get("gender") {
			return true
		}
	}
	//return false

	//复选框
	/*<input type="checkbox" name="interest" value="football">足球
	<input type="checkbox" name="interest" value="basketball">篮球
	<input type="checkbox" name="interest" value="tennis">网球*/
	slice2 := []string{"football", "basketball", "tennis"}
	a := Slice_diff(r.Form["interest"], slice2)
	if a == nil {
		return true
	}
	//return false

	//日期和时间
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	//身份证号码
	//验证15位身份证，15位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		return false
	}

	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		return false
	}
}

func Slice_diff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !In_slice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}
func In_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
