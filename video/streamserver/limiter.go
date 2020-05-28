package main

import (
	"fmt"
)

type ConnLimiter struct {
	// 规定的
	concurrentConn int
	// 动态变化的
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		return false
	}

	cl.bucket <- 1

	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	fmt.Print(c)
}
