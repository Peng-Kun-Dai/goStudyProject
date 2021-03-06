package main

import "fmt"

func sort(data sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	var sum IntArray
	sum = IntArray([]int{2, 3, 6, 4, 1}) //切片转型为复合类型
	fmt.Println(sum)
	sort(sum)
	fmt.Println(sum)

}
