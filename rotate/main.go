package main

import "fmt"

func main() {
	fmt.Println(leftRotate("abcde", 2))
	fmt.Println(rightRotate("abcde", 2))
}

func leftRotate(s string, shift int) string {
	var str string = ""
	for i := shift; i < len(s); i++ {
		str = str + string(s[i])
	}
	for i := 0; i < shift; i++ {
		str = str + string(s[i])
	}

	return str
}

func rightRotate(s string, shift int) string {
	var str string = ""
	for i := len(s) - shift; i < len(s); i++ {
		str = str + string(s[i])
	}
	for i := 0; i < len(s)-shift; i++ {
		str = str + string(s[i])
	}

	return str
}
