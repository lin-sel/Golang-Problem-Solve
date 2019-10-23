package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abbaeae"
	stp := []rune(str)
	count := 0
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			if j-i >= 2 {
				if strings.Compare(rev(string(stp[i:j])), string(stp[i:j])) == 0 {
					fmt.Println(string(stp[i:j]))
					count++
				}
			}
		}
	}
}

func rev(s1 string) string {

	s := []rune(s1)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
