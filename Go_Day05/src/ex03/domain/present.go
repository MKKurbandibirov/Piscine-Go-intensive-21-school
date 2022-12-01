package domain

import "fmt"

type Present struct {
	Value int
	Size  int
}

func (p *Present) String() string {
	return fmt.Sprintf("(%d, %d)", p.Value, p.Size)
}

func NewPresent(val, size int) *Present {
	return &Present{
		Value: val,
		Size:  size,
	}
}
