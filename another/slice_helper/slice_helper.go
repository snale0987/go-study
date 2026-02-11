package slice_helper

import (
	"fmt"
	"strings"

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

func slice_uniq_map() {
	type User struct {
		Name string
		Age  int
	}
	users := []User{{Name: "Alex", Age: 10}, {Name: "Alex", Age: 12}, {Name: "Bob", Age: 11}, {Name: "Alice", Age: 20}}
	uniqMap := lo.UniqMap(users, func(item User, index int) string {
		return item.Name
	})
	fmt.Println(uniqMap)
}

func slice_filter_map() {
	items := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(item string, index int) (string, bool) {
		if strings.HasSuffix(item, "pu") {
			return item, true
		}
		return "", false
	})
	fmt.Println(items)
}

func slice_flat_map() {
	type User struct {
		Name string
		Tags []string
	}
	users := []User{
		{Name: "张三", Tags: []string{"Go", "后端"}},
		{Name: "李四", Tags: []string{"Java", "后端"}},
		{Name: "王五", Tags: nil}, // 无标签
	}

	allTags := lo.FlatMap(users, func(u User, _ int) []string {
		return u.Tags
	})

	fmt.Println(allTags)

}
