package utils

import (
	"github.com/satori/go.uuid"
)

/**
 * @desc  生成uuid
 * @author Ipencil
 * @create 2018/12/9
 */

func NewUUID() (string,error) {
	// 创建
	u1, err := uuid.NewV4()
	return u1.String(),err
	//// 解析
	//u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	//if err != nil {
	//	fmt.Printf("Something gone wrong: %s", err)
	//	return ""
	//}
}
