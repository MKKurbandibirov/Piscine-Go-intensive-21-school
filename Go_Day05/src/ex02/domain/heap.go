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

type MaxHeap struct {
	Presents []*Present
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (m *MaxHeap) Len() int { return len(m.Presents) }
func (m *MaxHeap) Less(i, j int) bool {
	if m.Presents[i].Value > m.Presents[j].Value {
		return true
	} else if m.Presents[i].Value == m.Presents[j].Value && m.Presents[i].Size < m.Presents[j].Size {
		return true
	}
	return false
}
func (m *MaxHeap) Swap(i, j int) { m.Presents[i], m.Presents[j] = m.Presents[j], m.Presents[i] }

func (m *MaxHeap) Push(x interface{}) { m.Presents = append(m.Presents, x.(*Present)) }
func (m *MaxHeap) Pop() interface{} {
	res := m.Presents[len(m.Presents)-1]
	m.Presents = m.Presents[:len(m.Presents)-1]
	return res
}
