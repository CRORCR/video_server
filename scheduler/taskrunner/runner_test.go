package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

/**
 * @desc    测试生产者 消费者
 * @author Ipencil
 * @create 2018/12/22
 */
func TestRunner(t *testing.T) {
	//生产者
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("dispatcher send:%d", i)
		}
		return nil
	}
	//消费者
	e := func(dc dataChan) error {
		//这里用loop 意思是从dc取数据,取不到的时候,就退出,不要死循环等待
	loop:
		for {
			select {
			case d := <-dc:
				log.Printf("executor received:%d", d)
			default:
				break loop
			}
		}
		return errors.New("exit")//执行完成后,抛出异常,在 r.Error <- CLOSE 就会把close存入r.error中
	}
	rune := NewRunner(30, false, d, e)
	go rune.StartAll() //必须go 启动,否则后面time就不会执行了
	time.Sleep(3 * time.Second)
}
/*
for d:=range chan{

}
以上写法有坑,使用range是同步的,顺序执行,所以会阻塞,如果取完之后,没有close,就会一直等待

这里使用loop break方式
 */
