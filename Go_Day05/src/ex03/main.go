package main

import (
	"day05/domain"
	"math"
)

func findAns(res []int, presents []domain.Present, dp [][]int, i, s int) {
	if dp[i][s] == 0 {
		return
	}
	if dp[i-1][s] == dp[i][s] {
		findAns(res, presents, dp, i-1, s)
	} else {
		findAns(res, presents, dp, i-1, s-presents[i].Size)
		res = append(res, i)
	}
}

func grabPresents(presents []domain.Present, size int) []domain.Present {
	dp := make([][]int, len(presents)+1)
	for i := 1; i < len(presents)+1; i++ {
		dp[i] = append(dp[i], size+1)
		for s := 1; s < size; s++ {
			if s >= presents[i].Size {
				dp[i][s] = int(math.Max(float64(dp[i-1][s]), float64(dp[i-1][s-presents[i].Size]+presents[i].Value)))
			} else {
				dp[i][s] = dp[i-1][s]
			}
		}
	}

	ans := make([]int, 0)
	findAns(ans, presents, dp, len(presents)+1, size+1)
	println(ans)
	return nil
}

func main() {
	presents := make([]domain.Present, 0)
	presents = append(presents, *domain.NewPresent(5, 1))
	presents = append(presents, *domain.NewPresent(4, 5))
	presents = append(presents, *domain.NewPresent(3, 1))
	presents = append(presents, *domain.NewPresent(5, 2))
	grabPresents(presents, 6)
}
