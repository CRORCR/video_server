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
	loop:
		for {
			select {
			case d := <-dc:
				log.Printf("executor received:%d", d)
			default:
				break loop
			}
		}
		return errors.New("exit")
	}
	rune := NewRunner(30, false, d, e)
	go rune.StartAll()
	time.Sleep(3 * time.Second)
}
