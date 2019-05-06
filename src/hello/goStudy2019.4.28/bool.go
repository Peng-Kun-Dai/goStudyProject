package main

import "fmt"

func main() {
	var b bool
	//b=true
	b = (1 < 2)
	fmt.Println(b)

	s := newSlice()
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println(s[4])
	//p := &s
	s.remove(4)
}
func isEqual(i, j interface{}) bool {
	if i == j {
		return true
	} else {
		return false
	}
}

type Slice []int

func newSlice() Slice {
	return make([]int, 5)
}
func (s *Slice) remove(value interface{}) {
	for i, v := range *s {
		if isEqual(value, v) {
			*s = append((*s)[:i], (*s)[i+1:]...) //当i为最后一个元素是时，等效于追加一个空切片
			fmt.Println(*s)
			break
		}
	}

}
