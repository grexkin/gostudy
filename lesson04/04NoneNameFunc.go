package main

import "fmt"

func add(a,b int) int {
	return a + b
}

func testFunc1()  {
	f1 := add
	fmt.Printf("Type of f1: %T\n",f1)
	sum := f1(2,5)
	fmt.Printf("sum=%d\n",sum)
}

func testFunc2()  {
	f1 := func(a,b int) int {
		return a + b
	}
	fmt.Printf("type of f1=%T\n",f1)
	sum := f1(2,4)
	fmt.Printf("sum=%d\n",sum)
}

func testFunc3()  {
	var i int = 0
	defer fmt.Printf("defer i=%d\n",i)
	i = 100
	fmt.Printf("i=%d\n",i)
	return
}

func testFunc4()  {
	var i int = 0
	defer func() {
		fmt.Printf("defer i=%d\n",i)
	}()
	i = 100
	fmt.Printf("i=%d\n",i)
}

func calc(a,b int,op func(int,int)int) int {
	return op(a,b)
}

func add2(a,b int) int{
	return a + b
}

func sub(a,b int) int {
	return a - b
}

func testFunc5()  {
	sum := calc(100,200,add2)
	sub := calc(100,200,sub)
	fmt.Printf("sum=%d,sub=%d\n",sum,sub)
}

func main() {
	//testFunc1()
	//testFunc2()
	//testFunc3()
	//testFunc4()
	testFunc5()
}