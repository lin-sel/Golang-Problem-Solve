package main

import "fmt"

func main() {

	mp := make(map[int]int)

	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(findMissingNumber(mp, arr[:]))
}

func findMissingNumber(mp map[int]int, arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = mp[arr[i]] + 1
		if max < arr[i] {
			max = arr[i]
		}
	}

	if mp[1] == 0 {
		return 1
	} else if mp[2] == 0 {
		return 2
	}

	for i := 3; i <= max; i++ {
		if mp[i] == 0 {
			return i
		}
	}

	return max + 1
}
