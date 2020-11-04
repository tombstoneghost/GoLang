package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var increment int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			r := atomic.LoadInt64(&increment)
			fmt.Println(r)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Value:", increment)
}
