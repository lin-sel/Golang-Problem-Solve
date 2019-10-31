package main

import "fmt"

func main() {
	fmt.Println(LetterChange("abcd efg"))
	fmt.Println(LetterChange("azpqu!"))
}

// LetterChange shift string char to next one.
func LetterChange(str string) string {
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		var temp rune
		if arr[i] == 90 || arr[i] == 122 {
			temp = rune(int(str[i]) - 25)
		} else if arr[i] == ' ' || (int(arr[i]) >= 91 && int(arr[i]) < 97) || int(arr[i]) > 122 || int(arr[i]) < 65 {
			continue
		} else {
			temp = rune(int(str[i]) + 1)
		}

		if temp == 'a' || temp == 'i' || temp == 'e' || temp == 'o' || temp == 'u' {
			arr[i] = rune(ToUpper(int(temp)))
		} else {
			arr[i] = rune(temp)
		}
	}
	return string(arr)
}

// ToUpper change char to upper char
func ToUpper(val int) int {
	if val > 96 {
		val = val - 97
		val = val + 65
	}
	return val
}
