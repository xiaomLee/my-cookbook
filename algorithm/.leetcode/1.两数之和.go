package leetcode

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	quickSort(nums)
	fmt.Println(nums)
	n := len(nums)
	i := 0
	j := n - 1
	for i < j {
		math.MinInt32
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			break
		}
	}
	return []int{i, j}
}

func quickSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// @lc code=end
