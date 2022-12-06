package main

import (
	"day05/domain"
	"fmt"
	"math"
	"sort"
)

func findAns(res *[]int, presents []domain.Present, dp [][]int, i, s int) {
	if i < 0 || dp[i][s] == 0 {
		return
	}
	if i > 0 && dp[i-1][s] == dp[i][s] {
		findAns(res, presents, dp, i-1, s)
	} else {
		findAns(res, presents, dp, i-1, s-presents[i].Size)
		*res = append(*res, i)
	}
}

func grabPresents(presents []domain.Present, size int) []domain.Present {
	sort.Slice(presents, func(i, j int) bool {
		if presents[i].Size < presents[j].Size {
			return true
		}
		return presents[i].Size == presents[j].Size && presents[i].Value < presents[j].Value
	})
	dp := make([][]int, len(presents))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, size+1)
		for s := 1; s <= size; s++ {
			if s >= presents[i].Size {
				if i == 0 {
					dp[i][s] = presents[i].Value
				} else {
					dp[i][s] = int(math.Max(float64(dp[i-1][s]), float64(dp[i-1][s-presents[i].Size]+presents[i].Value)))
				}
			} else {
				dp[i][s] = dp[i-1][s]
			}
		}
	}
	ans := make([]int, 0)
	findAns(&ans, presents, dp, len(presents)-1, size)

	result := make([]domain.Present, 0)
	for i := 0; i < len(ans); i++ {
		result = append(result, presents[ans[i]])
	}
	return result
}

func main() {
	presents := make([]domain.Present, 0)
	presents = append(presents, *domain.NewPresent(5, 1))
	presents = append(presents, *domain.NewPresent(4, 5))
	presents = append(presents, *domain.NewPresent(3, 1))
	presents = append(presents, *domain.NewPresent(5, 2))
	result := grabPresents(presents, 6)
	fmt.Println(result)
}
