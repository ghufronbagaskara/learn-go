package main

import (
	"fmt"
	"sync"
)


func DemoRange()  {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i < 10; i++ {
		jobs <- i
	}

	close(jobs)
	
	wg.Wait()
}


func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker id: %d is processing job %d\n", id, job)
	}
}