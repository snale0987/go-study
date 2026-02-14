package slice_helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
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

func slice_reduce() {
	sum := lo.Reduce([]int{1, 2, 3, 4, 5}, func(acc int, item int, _ int) int {
		return acc + item
	}, 0)
	fmt.Println(sum)
}

func slice_reduce_right() {
	sum := lo.ReduceRight([]int{1, 2, 3, 4, 5}, func(acc int, item int, _ int) int {
		return acc + item
	}, 0)
	fmt.Println(sum)
}

func slice_foreach() {
	lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})
}

func slice_foreach_while() {
	list := []int64{1, 2, -42, 4}
	lo.ForEachWhile(list, func(x int64, _ int) bool {
		if x < 0 {
			return false
		}
		fmt.Println(x)
		return true
	})
}

func slice_times() {
	times := lo.Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Println(times)
}

func slice_parallel_times() {
	times := parallel.Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Println(times)
}

func slice_uniq() {
	uniqValues := lo.Uniq([]int{1, 2, 2, 1})
	fmt.Println(uniqValues)
}

func slice_uniq_by() {
	uniqByValues := lo.UniqBy([]int{1, 2, 3, 4, 5, 6}, func(i int) int {
		return i % 3
	})
	fmt.Println(uniqByValues)
}

func slice_group_by() {
	groupByValues := lo.GroupBy([]int{1, 2, 3, 4, 5, 6}, func(i int) int {
		return i % 3
	})
	fmt.Println(groupByValues)
}

func slice_group_by_map() {
	groupByMap := lo.GroupByMap([]int{1, 2, 3, 4, 5, 6}, func(i int) (int, int) {
		return i % 3, i * 2
	})
	fmt.Println(groupByMap)
}

func slice_chunk() {
	chunk1 := lo.Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	fmt.Println(chunk1)
	// [][]int{{0, 1}, {2, 3}, {4, 5}}

	chunk2 := lo.Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	// [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}
	fmt.Println(chunk2)

	chunk3 := lo.Chunk([]int{}, 2)
	// [][]int{}
	fmt.Println(chunk3)

	chunk4 := lo.Chunk([]int{0}, 2)
	// [][]int{{0}}
	fmt.Println(chunk4)
}

func slice_partitions() {
	partitions := lo.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println(partitions)
}
