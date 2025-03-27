package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	GoAdd()
	return
	var manorAttr map[int32]int32 = nil
	for k, v := range manorAttr {
		fmt.Println(k, v)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	var anyValue interface{}
	anyValue = "hello"
	anyValue = 123

	av := anyValue.(string)
	fmt.Println(av)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context is canceled, exiting...")
				wg.Done()
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	wg.Wait()
}

func GoAdd() {
	var allSum atomic.Int32 = atomic.Int32{}
	wg := sync.WaitGroup{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10000; i++ {
				allSum.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(allSum)
}

func GoAddInt() {
	var allSum int32 = 0
	wg := sync.WaitGroup{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10000; i++ {
				allSum++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(allSum)
}

func GoAddLock() {
	mutex := &sync.Mutex{}
	var allSum int32 = 0
	wg := sync.WaitGroup{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10000; i++ {
				mutex.Lock()
				allSum++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(allSum)
}

func GoAddBigLock() {
	mutex := &sync.Mutex{}
	var allSum int32 = 0
	wg := sync.WaitGroup{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			for i := 0; i < 10000; i++ {
				allSum++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(allSum)
}

func GoAddIntSerial() {
	var allSum int32 = 0
	for j := 0; j < 100; j++ {

		for i := 0; i < 10000; i++ {
			allSum++
		}
	}

	fmt.Println(allSum)
}
