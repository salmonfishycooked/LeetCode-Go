package _34_find_first_and_last_position_of_element_in_sorted_array

import (
	"fmt"
	"testing"
)

func TestProblem34(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Printf("Input: %v\nOutput: %v\n", nums, searchRange(nums, target))
}
