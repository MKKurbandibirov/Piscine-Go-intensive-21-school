package mean

type Mean float64

func ComputeMean(arr []int) Mean {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return Mean(float64(sum) / float64(len(arr)))
}
