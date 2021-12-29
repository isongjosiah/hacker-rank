package angry_professor

import (
	"sort"
)

func AngryProfessor(threshold int32, input []int32) string {
	sort.SliceStable(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	zi := getMinIndexes(input)
	if len(zi) == 0 {
		return "YES"
	}
	if len(zi) == 1 {
		onTime := int32(len(input[0:zi[0]]))
		if onTime >= threshold {
			return "NO"
		}
	}
	onTime := int32(len(input[0 : zi[len(zi)-1]+1]))
	if onTime >= threshold {
		return "NO"
	}
	return "YES"
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

	}

	return ind
}
