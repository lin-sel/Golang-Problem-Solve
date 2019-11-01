package main

import "fmt"

func main() {
	arr := []int{11, 22, 33, 44, 44, 66}
	mp := make(map[int]int)
	mp = frequentNumber(arr, mp)

	for i, j := range mp {
		fmt.Println(i, j)
	}
}

func frequentNumber(arr []int, mp map[int]int) map[int]int {

	for _, i := range arr {
		mp[i] = mp[i] + 1
	}

	return mp
}
