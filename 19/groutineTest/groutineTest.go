package main

import (
	"fmt"
	"runtime"
	"sync"
)

func log(msg interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Println(runtime.FuncForPC(pc).Name(), " ", msg)
}

func concurrencyTest() {
	res := 0
	m := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				m.Unlock()
				wg.Done()
			}()
			m.Lock()
			res++
		}()
	}
	wg.Wait()
	log(res)
}

func test1() {
	g := new(sync.WaitGroup)
	g.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer g.Done()
			log(i)
		}(i)
	}
	g.Wait()
}

func test2() {
	g := new(sync.WaitGroup)
	g.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer g.Done()
			log(i)
		}()
	}
	g.Wait()
}

func main() {
	test1()
	test2()
	concurrencyTest()
}
