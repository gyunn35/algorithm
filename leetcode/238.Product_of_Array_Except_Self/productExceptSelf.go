package leetcode

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := range res {
		res[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
