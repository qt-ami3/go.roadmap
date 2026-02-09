package main

import (
	"fmt"
	"sync"
)

func increment(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for i := 0; i < 1000000; i++ {
		mu.Lock()
		(*counter)++
		mu.Unlock()
	}
}

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)

	go increment(&counter, &wg, &mu)
	go increment(&counter, &wg, &mu)
	go increment(&counter, &wg, &mu)

	wg.Wait()

	fmt.Println("final counter:", counter)
}
