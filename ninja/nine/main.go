package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	sysinfo()

	wg.Add(1)
	go randomPrintloop("foo")

	wg.Add(1)
	go randomPrintloop("bar")

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

func randomPrintloop(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
	}
	wg.Done()
}
