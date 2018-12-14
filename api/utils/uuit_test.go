package utils

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/9
 */

func TestNewUUID(t *testing.T) {
	uuid, _ := NewUUID()
	fmt.Println(uuid)
}

func TestTime(t *testing.T) {
	ctime := time.Now().Format("Jan 02 2006,15:04:05")
	fmt.Println(ctime)
}
