// How to shift of bit using bitflags

package main

import  "fmt"

const (
	Readable = 1 << iota  	// 1 << 0 = 001
	Writable 				// 1 << 1 = 010
	Executable  			// 1 << 2 = 100
)


func main() {
	fmt.Println("\nFile permission:")
	fmt.Printf("Readable: %03b\n", Readable)
	fmt.Printf("Writable: %03b\n", Writable)
	fmt.Printf("Executable: %03b\n", Executable)
}