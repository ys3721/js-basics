package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math"
	"math/rand"
	"runtime"
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

func TestTwoWorkersAndOneConsumer(t *testing.T) {
	productsChan := make(chan int32)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	//defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)
	go worker1(ctx, productsChan, &wg)
	go worker2(ctx, productsChan, &wg)
	go consumer(ctx, productsChan, &wg)
	wg.Wait()
}

func worker1(ctx context.Context, productsChan chan<- int32, wg *sync.WaitGroup) {
	//每100毫秒生产一个产品
	defer wg.Done()
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker1 done")
			return
		case <-ticker.C:
			time.Sleep(100 * time.Millisecond)
			p := rand.Int31()
			productsChan <- p
			fmt.Printf("worker1 produce %d\n", p)
		}
	}
}

func worker2(ctx context.Context, productsChan chan<- int32, wg *sync.WaitGroup) {
	//每200毫秒生产一个产品
	defer wg.Done()
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker2 done")
			return
		case <-ticker.C:
			p := rand.Int31()
			productsChan <- p
			fmt.Printf("worker2 produce %d\n", p)
		}

	}
}

func consumer(ctx context.Context, productsChan <-chan int32, wg *sync.WaitGroup) {
	//每50毫秒消费一个产品
	defer wg.Done()
	ticker := time.NewTicker(50 * time.Millisecond)
	for {
		select {
		case <-ctx.Done(): //通道关闭从通道读数据不会报错 而是立刻返回一个=0值
			fmt.Println("consumer done")
			return
		case <-ticker.C:
			p := <-productsChan
			fmt.Printf("consumer consume %d\n", p)
		}
	}
}

func TestCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-ctx.Done(): //通道关闭从通道读数据不会报错 而是立刻返回一个=0值
				fmt.Println("ctx1 cancel done")
				return
			case <-ticker.C:
				fmt.Println("tick 1")
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		for {
			select {
			case <-ctx.Done(): //通道关闭从通道读数据不会报错 而是立刻返回一个=0值
				fmt.Println("ctx2 cancel done")
				return
			case <-ticker.C:
				fmt.Println("tick 2")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
}

func TestGoroutineSumNumber(t *testing.T) {
	//生成一个大的10000000个数的数组
	totalNumberCounts := 10000000
	numbers := make([]int32, totalNumberCounts)
	for i := 0; i < totalNumberCounts; i++ {
		numbers[i] = rand.Int31()
	}

}

func TestSpecialGoroutine(t *testing.T) {
	bes := splitHoleArrays(10, 3)
	fmt.Println(bes)
}
func splitArray(n int, numThreads int) {
	startIndex := 0
	chunkSize := n / numThreads
	remainder := n % numThreads

	for i := 0; i < numThreads; i++ {
		size := chunkSize
		if i < remainder {
			size++ // 前 remainder 个线程多拿一个元素
		}
		endIndex := startIndex + size - 1
		fmt.Printf("Thread %d: startIndex = %d, endIndex = %d\n", i, startIndex, endIndex)
		startIndex = endIndex + 1
	}
}

func splitArrays(n int, numThreads int) []BeginEnd {
	bes := make([]BeginEnd, numThreads)
	chuckSize := n / numThreads
	remainder := n % numThreads
	startIndex := 0
	for i := 0; i < numThreads; i++ {
		size := chuckSize
		if i < remainder {
			size++
		}
		endIndex := startIndex + size - 1
		be := BeginEnd{
			BeginIndex: startIndex,
			EndIndex:   endIndex,
		}
		bes[i] = be
		startIndex = endIndex + 1
	}
	return bes
}

type BeginEnd struct {
	BeginIndex int
	EndIndex   int
}

func splitHoleArrays(n int, numThreads int) []BeginEnd {
	chuckSize := n / numThreads
	remainder := n % numThreads
	segments := make([]BeginEnd, 0)

	startIndex := 0
	for i := 0; i < numThreads; i++ {
		endIndex := startIndex + chuckSize - 1
		segments = append(segments, BeginEnd{startIndex, endIndex})
		startIndex = endIndex + 1
	}
	if remainder > 0 {
		endIndex := startIndex + remainder
		segments = append(segments, BeginEnd{startIndex, endIndex})
	}
	return segments
}

func TestPrintAbc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	aChan := make(chan struct{})
	bChan := make(chan struct{})
	cChan := make(chan struct{})

	go func() { aChan <- struct{}{} }()
	go printTenTimes(ctx, "A", aChan, bChan)
	go printTenTimes(ctx, "B", bChan, cChan)
	go printTenTimes(ctx, "C", cChan, aChan)

	select {}
}

func printTenTimes(ctx context.Context, litter string, curChan <-chan struct{}, nextChan chan<- struct{}) {
	for {
		select {
		case <-curChan:
			fmt.Println(litter)
			nextChan <- struct{}{}
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}
}

func splitArray2(n int, numThreads int) [][]int {
	chuckSize := n / numThreads
	remainder := n % numThreads

	startIndex := 0

	segments := make([][]int, 0)

	startIndex = 0
	for i := 0; i < numThreads; i++ {
		endIndex := startIndex + chuckSize - 1
		if i < remainder {
			endIndex++
		}
		segments = append(segments, []int{startIndex, endIndex})
		startIndex = endIndex + 1
	}
	return segments
}

func TestSplitArray(t *testing.T) {
	segments := splitArray2(10, 3)
	fmt.Println(segments)
}

func splitArraysPageAddOne(n int32, numThreads int) [][]int32 {
	chuckSize := n / int32(numThreads)
	remainder := n % int32(numThreads)
	startIndex := int32(0)

	segments := make([][]int32, 0)
	for i := 0; i < numThreads; i++ {
		endIndex := startIndex + chuckSize - 1
		segments = append(segments, []int32{startIndex, endIndex})
		startIndex = endIndex + 1
	}
	if remainder > 0 {
		endIndex := startIndex + remainder
		segments = append(segments, []int32{startIndex, endIndex})
	}
	return segments
}

func TestSplitArraysPageAddOne(t *testing.T) {
	segments := splitArraysPageAddOne(10, 3)
	fmt.Println(segments)
}

func TestSplitArraysPageAddOne2(t *testing.T) {
	bigArrays := make([]int32, 10000000)
	for i := 0; i < 10000000; i++ {
		bigArrays[i] = rand.Int31()
	}
	c := GoroutineSumNumber(bigArrays)
	n := SumOneThreadOrdinary(bigArrays)
	fmt.Println(c)
	fmt.Println(n)
}

func SumOneThreadOrdinary(numbers []int32) int64 {
	total := int64(0)
	for i := 0; i < 10000000; i++ {
		total += int64(numbers[i])
	}
	return total
}

func GoroutineSumNumber(numbers []int32) int64 {
	//生成一个大的10000000个数的数组
	totalNumberCounts := 10000000
	goroutineNum := runtime.NumCPU() * 2
	segements := splitArrays(totalNumberCounts, goroutineNum)

	countChan := make(chan int, goroutineNum)

	for _, segement := range segements {
		seg := segement
		go func(seg BeginEnd) {
			sum := int64(0)
			for i := seg.BeginIndex; i <= seg.EndIndex; i++ {
				sum += int64(numbers[i])
			}
			countChan <- int(sum)
		}(seg)
	}
	totalSum := int64(0)
	for i := 0; i < len(segements); i++ {
		totalSum += int64(<-countChan)
	}
	return totalSum
}

func BenchmarkSumOneThreadOrdinary(b *testing.B) {
	numbers := generateNumbers(10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumOneThreadOrdinary(numbers)
	}
}

func BenchmarkSumAddUseThread(b *testing.B) {
	numbers := generateNumbers(100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GoroutineSumNumber(numbers)
	}
}

func generateNumbers(n int) []int32 {
	nums := make([]int32, n)
	for i := range nums {
		nums[i] = rand.Int31n(100)
	}
	return nums
}

func DynamicTimer(ctx context.Context, updateChan chan time.Duration) {
	tickerInterval := time.Second
	tickChan := make(chan time.Time)

	go func() {
		for {
			select {
			case <-time.After(tickerInterval):
				tickChan <- time.Now()
			case <-ctx.Done():
				fmt.Println("tick done!")
				return
			}
		}
	}()
	for {
		select {
		case newInterval := <-updateChan:
			tickerInterval = newInterval
			fmt.Println(fmt.Sprintf("tick interval 修改成了: %v", newInterval))
		case <-tickChan:
			fmt.Println("tick")
		case <-ctx.Done():
			fmt.Println("interval done!")
			return
		}
	}
}

func TestDynamicTimer(t *testing.T) {
	updateChan := make(chan time.Duration)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	go DynamicTimer(ctx, updateChan)

	time.Sleep(time.Second * 5)
	updateChan <- 3 * time.Second
	time.Sleep(time.Second * 60)

}

func TestDispatchedAndSender(t *testing.T) {
	msgChan := make(chan func(), math.MaxInt32)

	go dispatch(msgChan)

	for i := 0; i < 10000; i++ {
		plId := int32(i)
		go player(plId, msgChan)
	}

	select {}
}

func dispatch(c <-chan func()) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fn := <-c
			fn()
		}

	}
}

func player(id int32, c chan func()) {
	ticker := time.NewTicker(time.Millisecond*time.Duration(rand.Int31n(1000)) + 1)
	sqe := 0
	for {
		select {
		case <-ticker.C:
			c <- func() {
				fmt.Println(fmt.Sprintf("玩家：%d,消息数%d,线程数=%d", id, sqe, runtime.NumGoroutine()))
				sqe++
			}
		}
	}
}

func TestBeginTest(t *testing.T) {
	beginTest()
}
