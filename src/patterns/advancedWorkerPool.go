package patterns

import (
	"context"
	"fmt"
)

type Job struct {
	Id int
}

type Result struct {
	JobId int
	Err   error
}

func workerAdvanced(ctx context.Context, workerId int, jobs <-chan Job, result chan<- Result) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Println("Worker with id: ", workerId, " is serving job id: ", job.Id)
			// time.Sleep(500 * time.Millisecond)
			result <- Result{
				JobId: job.Id,
				Err:   nil,
			}
		}

	}
}

func AdvancedWorkerPoolDriver() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numWorker := 3

	jobs := make(chan Job, 5)
	results := make(chan Result, 5)
	for i := 1; i <= numWorker; i++ {
		go workerAdvanced(ctx, i, jobs, results)
	}

	// pushing jobs
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- Job{Id: i}
		}
		close(jobs)
	}()

	// fetching result from channel
	for i := 1; i <= 5; i++ {
		ans := <-results
		fmt.Println("Result processed with value: ", ans.JobId)
	}
	close(results)
}
