package taskrunner

/**
 * @desc   生产者 消费者
 * @author Ipencil
 * @create 2018/12/22
 */

type Runner struct {
	Controller controllChan //消息通道
	Error      controllChan
	Data       dataChan //数据通道
	DataSize   int
	LongLived  bool //长期存在的视频,不回收
	Dispatcher fn   //调度者
	Executor   fn   //执行者
}
//调度
func NewRunner(size int, longLived bool, d, e fn) *Runner {
	return &Runner{
		Controller: make(controllChan, 1),
		Error:      make(controllChan, 1),
		Data:       make(dataChan, size),
		DataSize:   size,
		LongLived:  longLived,
		Dispatcher: d,
		Executor:   e,
	}
}

//调度
func (r *Runner) StartDispatch() {
	//()  这个括号不能丢,丢了只是声明,不会调用啊
	defer func() {
		if !r.LongLived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_EXECUTE
				}
			}

			if c == READY_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

//启动之前需要存入一个消息,否则程序无法运行,阻塞
func(r *Runner)StartAll(){
	r.Controller<-READY_DISPATCH
	r.StartDispatch()
}