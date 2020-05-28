package task

import (
	"time"
)

type Worker struct {
	// 初始化一个ticker，接受从系统传递过来的时间点，达到我们的目标时间点后，能够触发想做的事
	// 跑多久 什么时候跑
	ticker *time.Ticker
	// 常驻任务  跑的谁
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		// 新建ticker
		ticker: time.NewTicker(interval * time.Second),
		// 新建任务
		runner: r,
	}
}
func (w *Worker) startWorker() {
	for {
		select {
		// ticker的channel每隔三秒通一次
		case <-w.ticker.C:
			// 任务的启动器
			go w.runner.StartAll()
		}
	}
}

// 总的部分
func Start() {
	// 任务：清除视频  为此有一个找删除数据的   一个进行删除数据的
	// 任务：第一个datachannel的长度  第二个是是否永久保持   第三个是生产者   第四个是消费者
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	// 第一个是定时器   第二个是任务
	// 我们这样理解：
	w := NewWorker(3, r)
	go w.startWorker()
}
