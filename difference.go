func findTheDifference(s string, t string) byte {
    var result byte = 0

    for i := 0; i < len(s); i++ {
        result ^= s[i]
        result ^= t[i]
    }

    result ^= t[len(t)-1]
    return result
}


// func findTheDifference(s string, t string) byte {
// 	first := []rune(s)
// 	second := []rune(t)

// 	// Count characters
// 	count := make(map[rune]int)
// 	for _, ch := range first {
// 		count[ch]++
// 	}
// 	for _, ch := range second {
// 		count[ch]--
// 		if count[ch] < 0 {
// 			return byte(ch)
// 		}
// 	}
// 	return 0
// }
