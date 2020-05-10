package main

import (
	"fmt"
	"time"
)

func getNow()  {
	now := time.Now()
	fmt.Printf("current time is %v\n",now)
}

func formatDate() {
	year,month,day,
	hour,minute,second := time.Now().Year(),time.Now().Month(),time.Now().Day(),
	                      time.Now().Hour(),time.Now().Minute(),time.Now().Second()

	fmt.Printf("%02d-%02d-%02d %02d-%02d-%02d\n",year,month,day,hour,minute,second)
}

func timeStamp () {
	timeStamps := time.Now().Unix()
	fmt.Printf("current timestamp is %d\n",timeStamps)
	// 时间戳转换成时间格式
	timeObj := time.Unix(timeStamps,0)
	year,month,day,
	hour,minute,second := timeObj.Year(),timeObj.Month(),timeObj.Day(),
	timeObj.Hour(),timeObj.Minute(),timeObj.Second()
	fmt.Printf("%02d-%02d-%02d %02d-%02d-%02d\n",year,month,day,hour,minute,second)
}


func timeStampTransToDateTime(timestamp int64) {
	timeObj := time.Unix(timestamp,0)
	year,month,day,
	hour,minute,second := timeObj.Year(),timeObj.Month(),timeObj.Day(),
		timeObj.Hour(),timeObj.Minute(),timeObj.Second()
	fmt.Printf("%02d-%02d-%02d %02d-%02d-%02d\n",year,month,day,hour,minute,second)
}

func processTask()  {
	fmt.Printf("do task\n")
}

func tickTest()  {
	//定时器练习
	ticker := time.Tick(time.Second*1)
	for i := range ticker {
		fmt.Printf("%v\n",i)
		processTask()
	}
}

func timeConst()  {
	//纳秒
	fmt.Printf("nano second %d\n",time.Nanosecond)
	//微秒
	fmt.Printf("micro second %d\n",time.Microsecond)
	//毫秒
	fmt.Printf("milli second %d\n",time.Millisecond)
	//秒
	fmt.Printf("second %d\n",time.Second)
	/*
	nano second 1
	micro second 1000
	milli second 1000000
	second 1000000000
	*/
}

func formatOriginal() {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Printf("%s\n",timeStr)
}


func main() {
	//getNow()
	//formatDate()
	//timeStamp()
	//timeStampTransToDateTime(1219021901)
	//tickTest()
	//timeConst()
	formatOriginal()
}
