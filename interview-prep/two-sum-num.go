package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 2, 7, 4}
	target := 6
	
	data := twoSum(nums, target)
	fmt.Printf("result %v \n", data)
}

// using hashmap to store numbers
func twoSum(nums[]int, target int) []int {
	m := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num

		fmt.Println("complement", m[complement])
		
		if idx, ok := m[complement]; ok {
			fmt.Printf("\nidx => %d & i => %d \n ", idx, i)
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}