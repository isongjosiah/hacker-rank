package min_max_sum

import (
	"fmt"
	"sort"
)

func MinMaxSum(arr []int32) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	min := sumArray(arr[0 : len(arr)-1])
	max := sumArray(arr[1:])
	fmt.Printf("%v %v\n", min, max)
}

func sumArray(arr []int32) int64 {
	sum := int64(0)
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	return sum
}
