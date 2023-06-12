package leetcode

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	b := strings.Builder{}

	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteString("|")
		b.WriteString(str)
	}

	return b.String()
}

func (codec *Codec) Decode(strs string) []string {
	res := []string{}

	for len(strs) != 0 {
		count := 0
		for _, char := range strs {
			if char == rune('|') {
				break
			}
			count++
		}

		length, _ := strconv.Atoi(strs[:count])

		res = append(res, strs[count+1:count+1+length])

		strs = strs[count+1+length:]
	}

	return res
}
