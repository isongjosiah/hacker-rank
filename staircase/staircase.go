package staircase

import (
	"fmt"
	"strings"
)

func Staircase(n int32) {
	var stairs = make([][]string, n)
	count := n - 1

	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			if j >= count {
				stairs[i] = append(stairs[i], "#")
			} else {
				stairs[i] = append(stairs[i], " ")
			}
		}
		count--
	}
	for x := int32(0); x < n; x++ {
		fmt.Println(strings.Replace(strings.Join(stairs[x], ","), ",", "", -1))
	}
}
