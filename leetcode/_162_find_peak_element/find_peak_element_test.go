package _162_find_peak_element

import (
	"fmt"
	"testing"
)

func TestProblem162(t *testing.T) {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Printf("Input: %v\nOutput: %v\n", nums, findPeakElement(nums))
}
