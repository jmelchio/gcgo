package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	sysinfo()

	wg.Add(1)
	go randomPrintloop("foo", 5)

	wg.Add(1)
	go randomPrintloop("bar", 5)

	raceCondition(true)
	goingNuclear()

	fmt.Println("That's all for Ninja level nine Folks !!")

	sysinfo()
	wg.Wait()
}

func sysinfo() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}

func randomPrintloop(name string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(name, ":", i)
	}
	wg.Done()
}

func raceCondition(useMutex bool) {
	counter := 0
	numRoutines := 100

	var raceWg sync.WaitGroup
	raceWg.Add(numRoutines)
	var mut sync.Mutex

	for i := 0; i < numRoutines; i++ {
		go func() {
			if useMutex {
				mut.Lock()
			}
			localCount := counter
			// runtime.Gosched()
			localCount++
			counter = localCount
			if useMutex {
				mut.Unlock()
			}
			raceWg.Done()
		}()
		fmt.Println("GoRoutines at loop:", i, "is", runtime.NumGoroutine())
	}
	raceWg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Endcount:", counter)
}

func goingNuclear() {
	var counter int64
	numRoutines := 100

	var raceWg sync.WaitGroup
	raceWg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:\t", atomic.LoadInt64(&counter))
			raceWg.Done()
		}()
		fmt.Println("GoRoutines at loop:", i, "is", runtime.NumGoroutine())
	}
	raceWg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Endcount:", counter)
}
