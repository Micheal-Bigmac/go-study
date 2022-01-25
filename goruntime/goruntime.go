package goruntime

import (
	"fmt"
	"sync"
	"time"
)

func Loop() {
	for i := 0; i < 11; i++ {
		time.Sleep(time.Microsecond * 2)
		fmt.Printf("%d", i)
	}
}
func Loop1() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d", i)
	}
}

var chanInt chan int = make(chan int, 10)
var timeout chan bool = make(chan bool)

//  发送数据到channel
func Send() {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3

	time.Sleep(time.Second * 2)
	timeout <- true

}

// 从channel 接受数据
func Receive() {
	/*num := <-chanInt
	fmt.Println("reciive ", num)

	num = <-chanInt
	fmt.Println("reciive ", num)

	num = <-chanInt
	fmt.Println("reciive ", num)*/

	for {
		select {
		case num := <-chanInt:
			fmt.Println("num : ", num)
		case num := <-timeout:
			fmt.Println("timeout : ", num)

		}
	}
}

var SWG sync.WaitGroup

//读取数据
func Read() {
	for i := 0; i < 3; i++ {
		SWG.Add(1)
	}
}

// 写入数据
func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		SWG.Done()
	}

}
