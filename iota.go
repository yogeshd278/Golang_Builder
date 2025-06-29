//  Use of iota

package main

import "fmt"

// const (
// 	Monday = iota + 1
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// 	Sunday
// )

const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Day of the week:")
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("Sunday:", Sunday)
}