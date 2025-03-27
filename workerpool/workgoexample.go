package main

import (
	"fmt"
	"time"
)

func niuMa(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("create niuMa", id, "started")
	for job := range jobs {
		//fmt.Println("niuMa", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("niuMa", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 100
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 200; w++ {
		go niuMa(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
