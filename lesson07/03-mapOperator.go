package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rangeMap()  {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int,1024)
	for i:=0;i < 128;i++ {
		key := fmt.Sprintf("stu%d",i)
		value := rand.Intn(1000)
		a[key] = value
	}

	for key,value := range  a {
		fmt.Printf("a[%s]=%d\n",key,value)
	}
}

func deleteMap()  {
	var a map[string]int = make(map[string]int,10)
	a["stu01"] = 1001
	a["stu02"] = 1002
	a["stu03"] = 1003
	fmt.Printf("before delete a=%#v\n",a)
	delete(a,"stu02")
	fmt.Printf("after delete a=%#v\n",a)
	// 删除所有记录
	for key,_ := range a {
		delete(a,key)
	}
	fmt.Printf("after delete a=%#v\n",a)
}

func main() {
	deleteMap()
}
