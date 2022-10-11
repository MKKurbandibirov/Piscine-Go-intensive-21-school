package median

import (
	"sort"
	"testing"
)

func TestComputeMedian1(t *testing.T) {
	got := ComputeMedian([]int{1})
	want := Median(1)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeMedian2(t *testing.T) {
	test := []int{1, -3, 5, -9, 23, -478, 990, 1892, 43, 21, -7, 89, 92, 784, -223, 11, -23, 4454, -12232, 8977}
	sort.Ints(test)
	got := ComputeMedian(test)
	want := Median(16)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeMedian3(t *testing.T) {
	got := ComputeMedian([]int{0, 0, 0, 0, 0})
	want := Median(0)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
