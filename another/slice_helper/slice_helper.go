package slice_helper

import (
	"fmt"

	"github.com/samber/lo"
)

func slice_filter() {
	even := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	fmt.Println(even)

	type IntSlice []int
	var mySlice IntSlice = []int{1, 2, 3}
	filtered := lo.Filter(mySlice, func(item int, i int) bool { return item > 1 })
	fmt.Printf("返回值类型：%T\n", filtered)
}

func slice_map() {
	ints := lo.Map([]int{1, 2, 3}, func(item int, index int) int { return item * 2 })
	fmt.Println(ints)
}
