package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

func BenchmarkGoAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAdd()
	}
}

func BenchmarkGoAddInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddInt()
	}
}

func BenchmarkGoAddLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddLock()
	}
}

func BenchmarkGoAddBigLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddBigLock()
	}
}

func BenchmarkGoAddIntSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddIntSerial()
	}
}

func TestAntFor(t *testing.T) {
	pool, _ := ants.NewPool(3)
	defer pool.Release()
	for i := 0; i < 1000; i++ {
		_ = pool.Submit(func() {
			println(i)
		})
	}
	time.Sleep(1 * time.Second)
}

func TestGoroutinePrintABAnts(t *testing.T) {
	var wg sync.WaitGroup
	aChan := make(chan struct{}, 1)
	bChan := make(chan struct{}, 1)
	cChan := make(chan struct{}, 1)

	pool, _ := ants.NewPool(3)
	defer pool.Release()

	aChan <- struct{}{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = pool.Submit(func() {
			defer wg.Done()
			<-aChan
			printAFunc()
			bChan <- struct{}{}
		})
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = pool.Submit(func() {
			defer wg.Done()
			<-bChan
			printBFunc()
			cChan <- struct{}{}
		})
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = pool.Submit(func() {
			defer wg.Done()
			<-cChan
			printCFunc()
			aChan <- struct{}{}
		})
	}
	wg.Wait()
	fmt.Println("finish")
}

func TestGoroutinePrintABC(t *testing.T) {
	aChan := make(chan struct{}, 100)
	bChan := make(chan struct{})
	cChan := make(chan struct{})
	aChan <- struct{}{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			<-aChan
			printAFunc()
			bChan <- struct{}{}
		}()
		go func() {
			defer wg.Done()
			<-bChan
			printBFunc()
			cChan <- struct{}{}
		}()
		go func() {
			defer wg.Done()
			<-cChan
			printCFunc()
			aChan <- struct{}{}
		}()
	}
	wg.Wait()
	//这个不想要但是提交了
}

var printAFunc = func() {
	println("A")
}
var printBFunc = func() {
	println("B")
}
var printCFunc = func() {
	println("C")
}
