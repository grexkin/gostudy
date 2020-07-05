package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func modifyMap()  {
	a := map[string]int {
		"steve": 1000,
		"jamine": 1001,
	}
	a["jack"] = 800
	fmt.Printf("joined a =%v\n",a)
	b := a
	b["jack"] = 1003
	fmt.Printf("after modify a=%v\n",a)
}

func sortMap()  {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int,1024)
	for i:=0;i<128;i++ {
		key := fmt.Sprintf("stu%d",i)
		value := rand.Intn(1000)
		a[key] = value
	}

	var keys []string = make([]string,0,128)

	for key,value := range a {
		fmt.Printf("map[%s]=%d\n",key,value)
		keys = append(keys,key)
	}
	sort.Strings(keys)

	for _,val := range keys {
		fmt.Printf("key:%s,value:%d\n",val,a[val])
	}
}

func sliceMap()  {
	var slices []map[string]int
	slices = make([]map[string]int,5,16)
	for index,val := range slices {
		fmt.Printf("slices[%d]=%v\n",index,val)
	}
	slices[0] = make(map[string]int,16)
	slices[0]["stu00"] = 1000
	slices[0]["stu01"] = 1001
	slices[0]["stu02"] = 1002
	fmt.Printf("slices=%v\n",slices)
	for index,val := range slices {
		fmt.Printf("slices[%d]=%v\n",index,val)
	}
}

func mapSlice()  {
	var m map[string][]int = make(map[string][]int,16)
	key := "stu01"
	value,ok := m[key]
	if !ok {
		m[key] = make([]int,0,16)
		value = m[key]
	}
	value = append(value,100)
	value = append(value,200)
	value = append(value,300)
	m[key] = value
	fmt.Printf("m=%v\n",m)
}

func main() {
	//modifyMap()
	//sortMap()
	//sliceMap()
	mapSlice()
}

/*
作业
1. 词频统计
2. 存储学生信息，通过id查询map[int][]map[string]string
*/