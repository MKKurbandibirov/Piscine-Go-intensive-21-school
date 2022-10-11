package sd

import (
	"day00/mean"
	"math"
)

type SD float64

func ComputeSD(arr []int, mean mean.Mean) SD {
	var tmp float64
	for _, v := range arr {
		tmp += math.Pow(float64(v)-float64(mean), 2.0)
	}
	return SD(math.Sqrt(tmp / float64(len(arr))))
}
