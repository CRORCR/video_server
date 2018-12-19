package main

import "log"

/**
 * @desc    流控
 * @author Ipencil
 * @create 2018/12/14
 */

type ConnLimiter struct {
	concurrentConn int  //上限
	bucket  chan int
}

//构造
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{cc, make(chan int, cc)}
}

//buff chan :就是有缓存的chan

//写入chan
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn { //这个桶存储数据固定
		return false
	}
	//桶未满 就存一个数据
	cl.bucket <- 1
	return true
}

//读取chan
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket  //当连接关闭,就取出一个数据
	log.Printf("remove bucket:%d", c)
	return
}
