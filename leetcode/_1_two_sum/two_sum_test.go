package _1_two_sum

import (
	"fmt"
	"testing"
)

func TestProblem1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Input: %v\n Output: %v\n", nums, twoSum(nums, target))
}
