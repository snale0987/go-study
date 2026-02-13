package slice_helper

import "testing"

func Test_slice_filter(t *testing.T) {
	slice_filter()
}

func Test_slice_map(t *testing.T) {
	slice_map()
}

func Test_slice_uniq_map(t *testing.T) {
	slice_uniq_map()
}

func Test_slice_filter_uniq_map(t *testing.T) {
	slice_filter_map()
}

func Test_slice_flat_map(t *testing.T) {
	slice_flat_map()
}

func Test_slice_reduce(t *testing.T) {
	slice_reduce()
}

func Test_slice_reduce_right(t *testing.T) {
	slice_reduce_right()
}

func Test_slice_foreach(t *testing.T) {
	slice_foreach()
}

func Test_slice_foreach_while(t *testing.T) {
	slice_foreach_while()
}

func Test_slice_times(t *testing.T) {
	slice_times()
}

func Test_slice_parallel_times(t *testing.T) {
	slice_parallel_times()
}
