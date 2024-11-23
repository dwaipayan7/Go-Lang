package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working\n", i)
	fmt.Printf("Worker %d end\n", i)
	wg.Done()
}

func main() {
	fmt.Println("Dwaipayan")

	var wg sync.WaitGroup
	// 3 worker
	for i := 0; i <= 3; i++ {
		wg.Add(1) // increment the weight group counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Workers task completed")
}
