package main
 
import (
	"fmt"
	"sync"
	"time"
)
 
var locker = new(sync.Mutex)
// func NewCond(l Locker) *Cond：用于创建条件，根据实际情况传入sync.Mutex或者sync.RWMutex的
// 指针，一定要是指针，否则会发生复制导致锁的失效
var cond = sync.NewCond(locker)
var waitgroup sync.WaitGroup
 

func aaaa(x int) {
	cond.L.Lock()
	cond.Wait()   			//有wait必须有lock
	fmt.Println("aaaa",x)
	time.Sleep(time.Second * 1)
	defer func() {
		cond.L.Unlock() 	//释放锁
		waitgroup.Done()	//减一个协程
	}()
}
 

 
func main() {
 
	for i := 0; i < 5; i++ {
		go aaaa(i)
		waitgroup.Add(1*2)
	}
	fmt.Println("Signal:")
	cond.Signal()			//前面的cond就是表示唤醒加锁的：加wait就是唤醒一个加锁的协程  不加wait会唤醒所有没堵塞的加锁协程。
	time.Sleep(time.Second * 5)
 
	fmt.Println("Broadcast:") 
	cond.Broadcast() 		//唤醒剩下堵塞的
 
	waitgroup.Wait()       //保证所有协程执行完毕
}
 