// Package stringutil contains utility functions for working with strings.
package stringutil

import "strings"

// Reverse return its argument sting reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseVersion2(s string) string {
	reverse := ""
	for _,v := range s {
		reverse = string(v) + reverse // minus - every time a brand new string needs to be allocated in memory - strings in GO are immutable. 
	}
	return reverse;
}

func ReverseVersion3(s string) string {
	var sb strings.Builder
	for i:= len(s)-1;i>=0;i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
