package main

import (
	"fmt"
	"time"
)

//格式化输出时间的两个例子
func timeFormat1 () {
	TimeStr := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("%s\n",TimeStr)
}

func testCost() {
	start := time.Now().UnixNano()
	for i:=0;i<10;i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	fmt.Printf("代码执行耗时%d微秒\n",(end-start)/1000)
}

func main() {
	timeFormat1()
	testCost()
}