package main

import "fmt"

func main() {
	fmt.Println(anagram("abcd", "adCb"))
}

func anagram(str1, str2 string) bool {
	var flag bool = false
	if len(str1) != len(str2) {
		flag = false
	} else {
		var arr [27]int
		for i := 0; i < len(str1); i++ {
			arr[str1[i]%26] = arr[str1[i]%26] + 1
			arr[str2[i]%26] = arr[str2[i]%26] + 1
		}

		for i := 1; i < len(arr); i++ {
			if arr[i] != 0 {
				if arr[i]%2 == 0 {
					flag = true
				} else {
					flag = false
					break
				}
			}
		}
	}
	return flag
}
