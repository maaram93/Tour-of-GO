package stringutil

import (
	"fmt"
	"strings"
)

// WordCount - returns a map with word count
func WordCount(s string) {
	r := make(map[string]int)
	sl := strings.Fields(s)
	for _, v := range sl {
		if r[v] >= 1 {
			r[v]++
		} else {
			r[v] = 1
		}
		fmt.Println(v, r[v])
	}
}
