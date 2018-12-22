package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server_1_5/scheduler/dbops"
)

/**
 * @desc    延迟删除
 * @author Ipencil
 * @create 2018/12/22
 */

//1.从数据库读取 需要删除的数据  存入dataChan
//2.从dataChan中读取数据,执行

//查出videoId 存入 dataChan中
func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) //一次查出3条记录
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

//从dataChan中读取数据,执行
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}  //线程安全,多个goroute读写也是没有问题
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				//删除本地视频文件
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				//删除数据库记录
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid) //vid必须在这里传入啊,作为参数给go route
			//在闭包内调用go route 会拿到瞬时状态,不会把状态保存,所以需要把状态作为参数传递到go route
		default:
			break forloop
		}
	}

	//城会玩,遍历map,出现error,就赋值给error  直接把这个err return
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil { //只要有一个error,就抛出
			return false
		}
		return true
	})

	return err
}

//删除文件
func deleteVideo(vid string) error {
	//删除本地视频
	err := os.Remove(VIDEO_PATH + vid)
	//判断是否文件不存在
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}