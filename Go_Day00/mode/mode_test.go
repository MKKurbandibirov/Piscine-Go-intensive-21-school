package mode

import (
	"sort"
	"testing"
)

func TestComputeMode1(t *testing.T) {
	got := ComputeMode([]int{1})
	want := Mode(1)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeMode2(t *testing.T) {
	test := []int{1, -3, 5, -9, 23, -478, 990, 1892, 43, 21, -7, 89, 92, 784, -223, 11, 23, 4454, -12232, 8977}
	sort.Ints(test)
	got := ComputeMode(test)
	want := Mode(23)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeMode3(t *testing.T) {
	got := ComputeMode([]int{0, 0, 1, 1})
	want := Mode(0)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
