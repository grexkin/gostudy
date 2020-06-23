package main

import (
	"fmt"
	"math/rand"
	"time"
)

func copyArray()  {
	var a [3]int = [3]int{1,2,3}
	b := a
	b[0]=100
	fmt.Println(a)
	fmt.Println(b)
}

func copyArray2()  {
	var a [3]int = [3]int{1,2,3}
	//指针类型可以修改
	modify(&a)
	fmt.Println(a)
}
func modify(b *[3]int)  {
	b[0] = 100
}

func sumArray(a [10]int) int {
	var sum int = 0
	for _,val := range a {
		sum += val
	}
	return sum
}

func testSumArray()  {
	rand.Seed(time.Now().UnixNano())   //初始化随机数种子
	var b [10]int
	for i:=0;i<len(b);i++ {
		b[i] = rand.Intn(1000)
		//b[i] = rand.Int()
	}
	sum := sumArray(b)
	fmt.Printf("sum=%d\n",sum)
}

func TwoSumArray(a [5]int,target int)  {
	for i:=0;i<len(a);i++ {
		other := target - a[i]
		for j := i+1;j<len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n",i,j)
			}
		}
	}
}

func testTwoSumArray()  {
	b := [...]int {1,3,5,8,7}
	TwoSumArray(b,8)
}
func main() {
	//copyArray()
	//copyArray2()
	//testSumArray()
	testTwoSumArray()
}
