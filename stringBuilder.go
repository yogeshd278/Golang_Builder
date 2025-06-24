package main

import (
	"fmt"
	"strings"
)


func display(word1 string, word2 string) {

	first := []rune(word1)
	second := []rune(word2)

	var builder strings.Builder

	minLen := len(first)
	if len(second) < minLen {
		minLen = len(second)
	}

	for index := 0; index < minLen; index++ {
		builder.WriteString(string(first[index]))
		builder.WriteString(string(second[index]))
	}

	if len(first) > minLen {
		builder.WriteString(string(first[minLen:]))
	}

	if len(second) > minLen {
		builder.WriteString(string(second[minLen:]))
	}

	result := builder.String()

	fmt.Println("result ", result)
}


func main() {
	display("abcfgtyuiolk", "pqr")
}