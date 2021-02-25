package main

import (
	"fmt"
	"sync"
)

func main() {
	var num = 0
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			num++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}
