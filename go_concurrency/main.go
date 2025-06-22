// main.go
package main

import (
	// "github.com/gin-gonic/gin"
	// "net/http"
	"time"
	"fmt"
	"sync"
)

func DisplayData(city string, ch chan<-string, wg *sync.WaitGroup) string {
	defer wg.Done()

	ch<-fmt.Sprintf("This is the %s", city)
	return city
}


func main() {

	startNow := time.Now()

	cities := []string{"Delhi", "Noida", "Gurgaon", "Faridabad", "Ghaziabad"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go DisplayData(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(startNow))
}