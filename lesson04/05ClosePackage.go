package main

import (
	"fmt"
	"strings"
	"time"
)

func adder() func(int) int  {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func testClose1()  {
	f := adder()
	ret := f(1)
	fmt.Printf("ret=%d\n",ret)
	ret2 := f(20)
	fmt.Printf("ret2=%d\n",ret2)
	ret3 := f(100)
	fmt.Printf("ret3=%d\n",ret3)

	f1 := adder()
	ret4 := f1(1)
	fmt.Printf("f1 ret4=%d\n",ret4)
}

func adder1(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func testClose2() {
	f := adder1(1)
	fmt.Printf("type of f:%#T, f=%d\n",f,f(1))
	fmt.Printf("f=%d\n",f(2))
}

func makeSufixFunc(sufix string) func(string) string{
	return func(name string) string {
		if !strings.HasSuffix(name,sufix) {
			return name+sufix
		}
		return name
	}
}

func testClose3() {
	f1 := makeSufixFunc(".bmp")
	f2 := makeSufixFunc(".png")
	bmp := f1("sss")
	bmp2 := f1("sss.bmp")
	png := f2("aaa")
	png2 := f2("aaa.png")
	fmt.Printf("bmp=%s,bmp2=%s\n",bmp,bmp2)
	fmt.Printf("png=%s,png2=%s\n",png,png2)
}


func calcs(base int)(func(int)int,func(int)int) {
	//fmt.Printf("base1=%d\n",base)
	add := func(i int) int{
		base += i
		//fmt.Printf("add1 base=%d\n",base)
		return base
	}
	//fmt.Printf("add2 base=%d\n",base)
	sub := func(i int) int {
		base -= i
		//fmt.Printf("sub1 base=%d\n",base)
		return base
	}
	//fmt.Printf("sub2 base=%d\n",base)
	return add,sub
}

func testClose4() {
	f1,f2 := calcs(10)
	fmt.Println(f1(1),f2(2))
	fmt.Println(f1(3),f2(4))
	fmt.Println(f1(5),f2(6))
}

func forClose()  {
	for i:=0;i<5;i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}

func main() {
	//testClose1()
	//testClose2()
	//testClose4()
	forClose()
}