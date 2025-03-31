package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func (t Task) Execute() {
	fmt.Printf("Task Execute %d\n", t.ID)
	time.Sleep(time.Millisecond * 500)
}

type Schedule struct {
	TaskQueue  chan Task
	WorkerPool chan struct{}
	wg         sync.WaitGroup
}

func NewSchedule(initWorkers int) Schedule {
	s := Schedule{
		TaskQueue:  make(chan Task, 100),
		WorkerPool: make(chan struct{}, initWorkers),
	}
	return s
}

func (s *Schedule) Start() {
	for i := 0; i < cap(s.WorkerPool); i++ {
		s.WorkerPool <- struct{}{} // 预填充初始 worker 数量
		go s.worker()
	}
}

func (s *Schedule) Stop() {
	close(s.TaskQueue)
	s.wg.Wait()
}

// 任务执行 worker
func (s *Schedule) worker() {
	s.wg.Add(1)
	defer s.wg.Done()
	<-s.WorkerPool
	for task := range s.TaskQueue {
		task.Execute()
	}
}

func (s *Schedule) ScaleUp(n int) {
	for i := 0; i < n; i++ {
		s.WorkerPool <- struct{}{}
		go s.worker()
	}
}

func beginTest() {
	scheduler := NewSchedule(2)
	scheduler.Start()

	go func() {
		for i := 0; i < 100; i++ {
			scheduler.TaskQueue <- Task{ID: i}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Expanding workers to 5")
	scheduler.ScaleUp(31)

	scheduler.wg.Wait()
	scheduler.Stop()
}
