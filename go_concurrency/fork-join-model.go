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

	time.Sleep(time.Second * 1) // that is know as fork-join model

	fmt.Println("End")
}