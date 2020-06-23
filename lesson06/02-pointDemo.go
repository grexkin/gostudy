package main

import "fmt"

func memAddr()  {
	var a int32
	a = 100
	fmt.Printf("%p,%v\n",&a,a)
}

func pointDemo() {
	var a int
	a = 100

	var b *int
	fmt.Printf("%p,%p,%v\n",b,&b,b)
	if b == nil {
		fmt.Printf("b is nill\n")
	}
	b = &a
	fmt.Printf("%v\n",b)
}

func pointDemo1()  {
	var a int
	a = 100
	var b *int = &a
	fmt.Printf("b=%d,addr b=%p\n",*b,b)
	*b = 200
	fmt.Printf("b=%d,a=%d,addr b=%p,addr a=%p\n",*b,a,b,&a)
	a ++
	fmt.Printf("b=%d,a=%d,addr b=%p,addr a=%p\n",*b,a,b,&a)
}

func modify(a *int)  {
	*a = 100
}

func testpoint()  {
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("b=%d,addr b=%p\n",b,&b)
}

func modify2(arr *[3]int)  {
	(*arr)[0] = 100
}

func testpoint2()  {
	var a [3]int = [3]int{1,2,3}
	modify2(&a)
	fmt.Printf("arr=%v\n",a)
}

func modifu3(sli []string)  {
	sli[0] = "aaa"
}

func testpoint3()  {
	s := []string{"a","b","c"}
	modifu3(s)
	fmt.Printf("s=%v\n",s)
}

func newPoint()  {
	var a *int = new(int)
	*a = 100
	fmt.Printf("a=%d\n",*a)
	fmt.Println("=====")
	var s *[]string = new([]string)    //分配内存
	fmt.Printf("s=%v,addr s=%p\n",s,&s)
	(*s) = make([]string,2,10)   //初始化
	(*s)[0] = "aaa"
	(*s)[1] = "bbb"
	fmt.Printf("s=%v\n",s)

}
func main() {
	//memAddr()
	//pointDemo()
	//pointDemo1()
	//testpoint()
	//testpoint2()
	//testpoint3()
	newPoint()
}