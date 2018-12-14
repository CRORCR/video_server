package main

import (
	"fmt"

)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2018/12/14
 */

 type Inter interface {
 	Name() string
 }

 type Student struct {
 	Names string
 	Age int
 }

 func(s *Student)Name()string{
 	fmt.Println("hello")
 	return "hello"
 }

func main() {
	handler := NewMiddleWareHandler()
	fmt.Println(handler.Name())
}

func NewMiddleWareHandler() Inter {
	s:=&Student{}
	s.Names="vvvv"
	return s
}