package flatland_space_stations

import "math"

func FlatLandSpaceStations(n int32, c []int32) int32 {
	if int32(len(c)) == n {
		return 0
	}

	var max = make([]int32, 0)
	for x := int32(0); x < n; x++ {
		tempMax := n
		for _, v := range c {
			diff := math.Abs(float64(v - x))
			if int32(diff) < tempMax {
				tempMax = int32(diff)
			}
		}
		max = append(max, tempMax)
	}
	i := max[0]
	for _, x := range max {
		if x > i {
			i = x
		}
	}
	return i
}
