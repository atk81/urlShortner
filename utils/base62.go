package utils

import (
	"bytes"
)

func Base62(n int) string {
	if n <= 0 {
		return "0"
	}
	var b62 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	var s []byte
	for n > 0 {
		s = append(s, b62[n%62])
		n /= 62
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func Base10(s string) int {
	if s == "" {
		return 0
	}
	var b62 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	var n int
	var base int = 1
	for i := len(s) - 1; i >= 0; i-- {
		n +=base*bytes.IndexByte(b62,s[i])
		base*=62
	}
	return n
}