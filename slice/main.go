package main

import (
	"fmt"
	"strings"
)

func main() {
	// Slice
	var str []rune
	str = append(str, 'A')
	fmt.Println(str)
	sliceManipulation("Nilesh")
}

func sliceManipulation(slice string) {
	runes := strings.Split(slice, "")
	fmt.Printf("%T", runes)

}
