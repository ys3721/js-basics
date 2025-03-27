package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID  int
	Msg string
}

func worker(id int, tasks <-chan *Task, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker id %d processing task %d with message %s\n", id, task.ID, task.Msg)
		time.Sleep(1 * time.Second)
		result <- fmt.Sprintf("Task %d processed by worker %d", task.ID, id)
	}
}

func main() {
	taskCount := 10
	workerCount := 3

	tasks := make(chan *Task, taskCount)
	results := make(chan string, taskCount)

	wg := &sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, tasks, results, wg)
	}

	for i := 0; i < taskCount; i++ {
		tasks <- &Task{
			ID:  i,
			Msg: fmt.Sprintf("Task %d", i),
		}
	}
	close(tasks)
	wg.Wait()
	close(results)
	for result := range results {
		fmt.Println(result)
	}
}
