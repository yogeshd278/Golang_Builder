package main

import "fmt"

func main() {
	fmt.Println("\t Print prefix sum array\n")

	arr := []int{2, 4, 1, 3, 6}

	pref := prefix(arr)
	fmt.Println(pref)
	
	// Sum of index 1 to 3

	l, r := 1, 3
	sum := pref[r]
	if l > 0 {
		sum = sum - pref[l - 1]
	}
	fmt.Println("\n", sum)
}


func prefix(arr []int) []int {

	pref := make([]int, len(arr))

	pref[0] = arr[0]

	for i:=1 ; i<len(arr); i++ {
		pref[i] = pref[i - 1] + arr[i]
	}
	return pref
}