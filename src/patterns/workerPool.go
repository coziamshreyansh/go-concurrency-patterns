package patterns

import (
	"fmt"
	"time"
)

func worker(w int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker ", w, " initiated for job ", j)
		time.Sleep(1 * time.Second)
		results <- j * 2
	}
}

func Driver() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 5; r++ {
		fmt.Println("Result was :", <-results)
	}
	close(results)
}
