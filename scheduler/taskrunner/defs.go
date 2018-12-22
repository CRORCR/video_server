package taskrunner

/**
 * @desc    定义初始化一些数据
 * @author Ipencil
 * @create 2018/12/22
 */

const (
	READY_DISPATCH="d" //调度
	READY_EXECUTE="e" //执行
	CLOSE="c"
	VIDEO_PATH ="K:\\upload\\video\\"
)
 type controllChan chan string //消息通道chan 存储d e c

 type dataChan chan interface{}  //数据chan

 type fn func(dc dataChan)error