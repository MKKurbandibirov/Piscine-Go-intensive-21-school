package mean

import "testing"

func TestComputeMean1(t *testing.T) {
	got := ComputeMean([]int{1})
	want := Mean(1)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeMean2(t *testing.T) {
	got := ComputeMean([]int{1, -3, 5, -9, 23, -478, 990, 1892, 43, 21, -7, 89, 92, 784, -223, 11, -23, 4454, -12232, 8977})
	want := Mean(220.35)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestComputeMean3(t *testing.T) {
	got := ComputeMean([]int{0, 0, 0, 0, 0})
	want := Mean(0)
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
