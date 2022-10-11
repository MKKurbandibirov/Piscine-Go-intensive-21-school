package sd

import (
	"day00/mean"
	"sort"
	"testing"
)

func TestComputeSD1(t *testing.T) {
	got := ComputeSD([]int{1}, mean.ComputeMean([]int{1}))
	want := SD(0)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeSD2(t *testing.T) {
	test := []int{1, -3, 5, -9, 23, -478, 990, 1892, 43, 21, -7, 89, 92, 784, -223}
	sort.Ints(test)
	got := ComputeSD(test, mean.ComputeMean(test))
	want := SD(564.5514640451795)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeSD3(t *testing.T) {
	got := ComputeSD([]int{0, 0, 1, 1}, mean.ComputeMean([]int{0, 0, 1, 1}))
	want := SD(0.5)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
