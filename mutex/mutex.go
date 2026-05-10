// Solve race condition

package main

import (
	"fmt"
	"sync"
)
type post struct {
	view int
	mu sync.Mutex
}


func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done() 
	}()
	
	p.mu.Lock()
	p.view += 1
}


func main() {

	var wg sync.WaitGroup

	myPost := post{view:0}

	for _ = range 100 {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.view)
}