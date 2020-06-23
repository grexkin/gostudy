package main

import (
	"fmt"
)

func testSlice()  {
	var a []int = make([]int,10)
	if a == nil {
		fmt.Printf("a is nil\n")
	}else  {
		fmt.Printf("a is %v\n",a)
	}
	a[0] = 100
}

func testSlice2()  {
	a := [5]int{1,2,3,4,5}
	var b []int
	b = a[1:5]
	fmt.Printf("b=%v\n",b)
	fmt.Printf("b[0]=%d\n",b[0])
	fmt.Printf("len(b)=%d,cap(b)=%d\n",len(b),cap(b))
	b=append(b,a[1:4]...)
	fmt.Printf("new b=%v\n",b)
	fmt.Printf("len(b)=%d,cap(b)=%d,Type(b)=%T\n",len(b),cap(b),b)
}

func testSlice3()  {
	a := [5]int{1,2,3,4,5}
	var b []int
	b = a[1:4]
	fmt.Printf("b=%v\n",b)
	c := a[1:]
	fmt.Printf("c=%v\n",c)
	d := a[:3]
	fmt.Printf("d=%v\n",d)
	e := a[:]
	fmt.Printf("e=%v\n",e)
}

func testSlice4()  {
	//切片的修改
	a := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("a=%v,type(a)=%T\n",a,a)
	b := a[2:5]
	fmt.Printf("b=%v,type(b)=%T\n",b,b)
	b[0] = b[0]+10
	b[1] = b[1]+20
	b[2] = b[2]+30
	for index,val := range a{
		fmt.Printf("a[%d]=%d\n",index,val)
	}
	for i := range a {
		a[i] = a[i]+10
	}
	fmt.Printf("after modify ,array a=%v,type(a)=%T\n",a,a)
}

func testSlice5()  {
	//切片是数组的引用
	a := [3]int{1,2,3}
	fmt.Printf("a=%v.type(a)=%T\n",a,a)
	b := a[:]
	c := a[:]
	b[0] = 100
	fmt.Printf("a=%v,b=%v\n",a,b)
	c[1] = 200
	fmt.Printf("a=%v,b=%v,c=%v\n",a,b,c)
}

func makeSlice()  {
	var a []int
	a = make([]int,5,10)
	a[0]=10
	fmt.Printf("a=%v，addr(a)=%p,len(a)=%d,cap(a)=%d\n",a,a,len(a),cap(a))
	a = append(a,11)
	fmt.Printf("a=%v，addr(a)=%p,len(a)=%d,cap(a)=%d\n",a,a,len(a),cap(a))
	for i:=0;i<8;i++ {
		a = append(a,i)
		fmt.Printf("a=%v，addr(a)=%p,len(a)=%d,cap(a)=%d\n",a,a,len(a),cap(a))
	}
	a = append(a,100)
	fmt.Printf("a=%v，addr(a)=%p,len(a)=%d,cap(a)=%d\n",a,a,len(a),cap(a))
}

func makeSlice2()  {
	a := [...]string{"a","b","c","d","e","f","g","h"}
	b := a[1:3]
	fmt.Printf("b=%v,len(b)=%d,cap(b)=%d\n",b,len(b),cap(b))
	c := b[:cap(b)]   //切片再切片
	fmt.Printf("c=%v,cap(c)=%d,len(c)=%d\n",c,cap(c),len(c))
}

func noneSlice()  {
	var a []int
	fmt.Printf("a=%v,len(a)=%d,cap(a)=%d\n",a,len(a),cap(a))
	if a == nil {
		fmt.Printf("a is nil\n")
	}
	a = append(a,100)
	a = append(a,200)
	for i:=0;i<1000;i++ {
		a = append(a,i)
	}
	fmt.Printf("a=%v,len(a)=%d,cap(a)=%d\n",a,len(a),cap(a))
}
func main() {
	//testSlice()
	//testSlice2()
	//testSlice3()
	//testSlice4()
	//testSlice5()
	//makeSlice2()
	noneSlice()
}
