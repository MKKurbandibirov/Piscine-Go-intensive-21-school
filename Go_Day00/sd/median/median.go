package median

type Median float64

func ComputeMedian(arr []int) Median {
	if len(arr)%2 == 0 {
		return Median(float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2)
	} else {
		return Median(float64(arr[len(arr)/2]))
	}
}
