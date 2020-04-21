package main

import (
	"fmt"
	"time"
)

type Notify interface{
	notify(chan<- int)
}
type Notifier struct {

}
//该函数就是限制对于通道c只能进行写入操作
func(n *Notifier)notify(c chan<- int){
	c <- 2
}

func testChan(){
	ch := make(chan int,1)
	nt :=  Notifier{}
	nt.notify(ch)
	select {
	case value,_ := <- ch:
		fmt.Printf("Receive notify:%+v",value)
	}
}


func testForTimer(c chan int){
	time.AfterFunc(2 * time.Second, func(){
		close(c)
	})
}
//函数会在2s后通道关闭后退出
func testFor(){
	ch := make(chan int, 1)
	testForTimer(ch)
	for value := range ch{
		fmt.Println(value)
	}
}
func main(){
	testFor()
}
