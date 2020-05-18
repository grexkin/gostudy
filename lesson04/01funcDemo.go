package main

import "fmt"

func testFunc01()  {
	fmt.Printf("hello none return func\n")
}

func AddFunc(a,b int) int {
	return a+b
}

func mutiReturns(a,b int) (sum,sub int) {
	sum = a + b
	sub = a - b
	return
}

func changeAbleArgs(a ...int) int{
	sum := 0
	for i:=0;i<len(a);i++ {
		sum = sum + a[i]
	}
	return sum
}

func changeAbleArgsV2 (a int,b ...int) int {
	//至少有一个参数
	sum := a
	for i:=0;i <len(b);i ++ {
		sum = sum + b[i]
	}
	return sum
}

func deferTest() {
	//defer 是先进后出的，适用于资源的释放，比如文件句柄，数据库连接的关闭
	defer fmt.Printf("defer1\n")
	fmt.Printf("hello\n")
	defer fmt.Printf("defer2\n")
}

func testDefer2()  {
	for i:=0;i<5;i++ {
		defer fmt.Printf("i=%d\n",i)
	}
	fmt.Printf("running\n")
	fmt.Printf("return\n")
}

func testDefer3()  {
	var i int = 0
	defer fmt.Printf("defer i=%d\n",i)   //输出0，定义之后的修改不会影响其值
	i = 1000
	fmt.Printf("run i=%d\n",i)
}

func main() {
	testFunc01()
	res := AddFunc(100,200)  //函数调用
	fmt.Printf("sum is %d\n",res)
	sum,sub := mutiReturns(100,200)
	fmt.Printf("sum=%d,sub=%d\n",sum,sub)
	sum1 := changeAbleArgs(1)  //可变参数是个切片，可以传0个或者多个参数
	fmt.Printf("sum1=%d\n",sum1)
	sum2 := changeAbleArgsV2(1,2,3,4)
	fmt.Printf("sum2=%d\n",sum2)
	deferTest()
	testDefer2()
	testDefer3()
}