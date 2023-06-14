package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp < 0 {
				j++
			} else if tmp > 0 {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j, k = j+1, k-1
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return res
}
