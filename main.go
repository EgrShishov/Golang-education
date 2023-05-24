package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding() int64 {
	var s1, s2 string
	var clear1, clear2 string
	fmt.Scan(&s1, &s2)

	for _, char := range s1 {
		if unicode.IsDigit(char) {
			clear1 += string(char)
		}
	}
	ans1, _err1 := strconv.ParseInt(clear1, 10, 64)

	for _, char := range s2 {
		if unicode.IsDigit(char) {
			clear2 += string(char)
		}
	}

	ans2, _err2 := strconv.ParseInt(clear2, 10, 64)

	if _err1 == nil && _err2 == nil {
		return ans2 + ans1
	}
	return 0
}

func main() {
	fmt.Print(adding())
}
