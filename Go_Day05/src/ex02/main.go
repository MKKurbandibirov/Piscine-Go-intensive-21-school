package main

import (
	"container/heap"
	"day05/domain"
	"fmt"
)

func getNCoolestPresents(presents []*domain.Present, n int) []*domain.Present {
	presentsHeap := domain.NewMaxHeap()
	heap.Init(presentsHeap)

	for i := 0; i < len(presents); i++ {
		heap.Push(presentsHeap, domain.NewPresent(presents[i].Value, presents[i].Size))
	}

	results := make([]*domain.Present, 0)
	for i := 0; i < n; i++ {
		coolest := heap.Pop(presentsHeap)
		results = append(results, coolest.(*domain.Present))
	}
	return results
}

func main() {
	presents := make([]*domain.Present, 0)
	presents = append(presents, domain.NewPresent(5, 1))
	presents = append(presents, domain.NewPresent(4, 5))
	presents = append(presents, domain.NewPresent(3, 1))
	presents = append(presents, domain.NewPresent(5, 2))
	fmt.Println(getNCoolestPresents(presents, 2))
}
