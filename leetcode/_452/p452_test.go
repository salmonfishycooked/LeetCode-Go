package _452

import (
	"fmt"
	"testing"
)

func TestProblem452(t *testing.T) {
	points := [][]int{[]int{3, 9}, []int{7, 12}, []int{3, 8}, []int{6, 8},
		[]int{9, 10}, []int{2, 9}, []int{0, 9}, []int{3, 9}, []int{0, 6}, []int{2, 8}}
	fmt.Println(findMinArrowShots(points))
}
