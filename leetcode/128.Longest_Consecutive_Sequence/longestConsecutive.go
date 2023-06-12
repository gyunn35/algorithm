package leetcode

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, n := range nums {
		m[n] = true
	}

	res := 0

	for n := range m {
		if !m[n-1] {
			seq := 1
			for m[n+seq] {
				seq++
			}

			if res < seq {
				res = seq
			}
		}
	}

	return res
}
