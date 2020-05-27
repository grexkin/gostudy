package main

import "fmt"

func defineArray()  {
	var a [3]int
	a = [3]int{1,2,3}
	fmt.Println(a)
	a[0] = 100
	fmt.Println(a)

	var b [4]int
	b[0] = 3
	b[1] = 4
	for i,j := range b {
		fmt.Println(i,j)
	}

	c := [...]int{1,2,3,4,5,6}
	fmt.Println(c,len(c))

	//对下标元素初始化
	d := [5]int{3:100,4:200}
	fmt.Println(d)
}

func defineArray2(){
	var a [2]int = [2]int{3,4}
	var b [2]int = [2]int{1,2}
	b = a
	fmt.Println(b)
}

func loopArray()  {
	var a [5]int = [5]int{2,3,4,3:100,4:200}
	for i:=0;i<len(a);i++ {
		fmt.Println(a[i])
	}
}
func main() {
	defineArray()
	defineArray2()
	loopArray()
}