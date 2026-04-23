package patterns

import (
	"fmt"
	"time"
)

func worker(wId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Job: ", j, " is being processed by worker: ", wId)
		time.Sleep(1 * time.Second)
		results <- j * 2
	}
}

func WorkerPoolDriver() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for j := 1; j <= 3; j++ {
		go worker(j, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println("Result was: ", <-results)
	}
	close(results)
}
