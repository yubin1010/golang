package main

import (
	"fmt"
	"sync"
	"time"
)

func dostuff(wg *sync.WaitGroup, i int) {
	fmt.Printf("goroutine id %d\n", i)
	//time.Sleep(60 * time.Second)
	time.Sleep(time.Second)
	fmt.Printf("goroutine2 id %d\n", i)
	wg.Done() // go routine이 끝이다
}

func main() {
	var wg sync.WaitGroup
	workers := 10

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go dostuff(&wg, i)
	}
	wg.Wait()
}
