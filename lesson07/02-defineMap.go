package main

import "fmt"

func demoMap01()  {
	var a map[string]int
	fmt.Printf("a=%v\n",a)
	if a == nil {
		//初始化
		a = make(map[string]int,1000)
		a["stu01"] = 101
	}
	//数据插入
	a["stu02"] = 102
	a["stu03"] = 103
	a["stu04"] = 104
	fmt.Printf("a=%#v\n",a)
}

func demoMap02()  {
	//声明时初始化map
	var a map[string]int = map[string]int{
		"stu01":101,
		"stu02":102,
		"stu03":103,
	}
	fmt.Printf("a=%#v\n",a)
	a["stu02"] = 1002    //更新操作
	fmt.Printf("a=%#v\n",a)
	// 通过key访问元素
	var key string = "stu03"
	fmt.Printf("stu03=%v\n",a[key])
	//判断元素是否存在
	value,ok := a["stu04"]
	if !ok {
		fmt.Printf("value of a[stu04] does bot exists\n")
		value = 1004
	}
	fmt.Printf("a[stu04]=%v\n",value)

	//map的遍历
	for k,v := range a {
		fmt.Printf("key=%s,value=%d\n",k,v)
	}
}

func main() {
	demoMap01()
	demoMap02()
}