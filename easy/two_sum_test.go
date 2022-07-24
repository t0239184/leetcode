package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Constraints:
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.
func Test_Two_Sum_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	expected := []int{0, 1}
	assert.Equal(t, expected, result)
}

func Test_Two_Sum_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	result := TwoSum(nums, target)
	expected := []int{1, 2}
	assert.Equal(t, expected, result)
}

func Test_Two_Sum_3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	result := TwoSum(nums, target)
	expected := []int{0, 1}
	assert.Equal(t, expected, result)
}
