package main

import (
	"fmt"
	"sync"
)

var locker = new(sync.Mutex)

var cond = sync.NewCond(locker)
// WaitGroup控制节奏的
var waitgroup sync.WaitGroup

func test(x int) {
	cond.L.Lock()
	fmt.Println("test", x)
	defer func() {
		cond.L.Unlock()  //释放锁
		waitgroup.Done() //减一个协程
	}()
}

func main() {

	for i := 0; i < 5; i++ {
		go test(i)
		waitgroup.Add(1)
	}

	waitgroup.Wait() //保证所有协程执行完毕
}

// test 0
// test 4
// test 3
// test 1
// test 2