package taskrunner

import (
	"time"
	//"log"
)

type Worker struct {
	ticker *time.Ticker //定时器
	runner *Runner
}

//duration是int64  但是这个time包自定义的
func NewWorker(t time.Duration, r *Runner) *Worker {
	return &Worker {
		ticker: time.NewTicker(t * time.Second),
		runner: r,
	}
}


func (w *Worker) startWorker() {
	for {
		select {
		case <- w.ticker.C:  //到期就执行
			go w.runner.StartAll()
		}
	}
}

func Start() {
	// Start video file cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}

/*
for c:= range w.ticker.C{}  这种方式慎用
<-w.ticker.C 使用这种方式
 */