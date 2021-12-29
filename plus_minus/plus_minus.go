package plus_minus

import (
	"fmt"
	"sort"
)

func PlusMinus(arr []int32) {
	var neg float32
	var pos float32
	var zero float32
	var total = float32(len(arr))

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	mi := getMinIndexes(arr)
	fmt.Println(arr)
	fmt.Println(mi, "minimum index")
	l := len(mi)
	if l == 1 {
		// check if minimum is zero
		if arr[mi[0]] == 0 {
			neg = float32(len(arr[0:mi[0]]))
			pos = float32(len(arr[mi[0]+1:]))
			zero = float32(l)
		} else {
			neg = float32(len(arr[0 : mi[0]+1]))
			pos = float32(len(arr[mi[0]+1:]))
			zero = 0
		}
	} else if l > 1 {
		// then it can only be a zero
		neg = float32(len(arr[0:mi[0]]))
		pos = float32(len(arr[mi[l-1]+1:]))
		zero = float32(l)
	} else if l == 0 {
		neg = 0
		pos = total
		zero = 0
	}

	negRatio := neg / total
	posRatio := pos / total
	zeroRatio := zero / total

	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}

func getMinIndexes(input []int32) []int {
	var ind []int
	for i, v := range input {
		if v == 0 {
			ind = append(ind, i)
		}
	}
	if len(ind) == 0 {
		//No zero in array. Get index of the last negative value
		ind = append(ind, 0)
		for i, v := range input {
			if v < 0 {
				ind[0] = i
			}
		}

		if ind[0] == 0 {
			// then there was no number less than zero
			return []int{}
		}

	}

	return ind
}
