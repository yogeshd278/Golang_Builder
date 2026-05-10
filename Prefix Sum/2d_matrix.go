package main

import "fmt"


func buildPrefixSum(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	prefix := make([][]int, rows)

	fmt.Println("prefix ", prefix)

	for i := range prefix {
		prefix[i] = make([]int, cols)
	}

	fmt.Println("\nprefix ", prefix)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			prefix[i][j] = mat[i][j]

			if i > 0 {
				prefix[i][j] += prefix[i-1][j]
			}
			if j > 0 {
				prefix[i][j] += prefix[i][j-1]
			}
			if i > 0 && j > 0 {
				prefix[i][j] -= prefix[i-1][j-1]
			}
		}
	}
	return prefix
}


func main() {

	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result := buildPrefixSum(mat)
	fmt.Println("\n\n",result)
}