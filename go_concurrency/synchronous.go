package main

import (
	"fmt"
	"time"
)

func Some(data string) {
	fmt.Println(data)
}

func main() {

	go Some("1")
	
	time.Sleep(time.Second * 2)
	
	go Some("2")
	go Some("3")



	fmt.Println("Hello")
}