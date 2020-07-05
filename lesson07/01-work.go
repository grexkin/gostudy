package main

import "fmt"

func modifyInt(a *int)  {
	*a = 1000
}

func swich(a,b *int)  {
	fmt.Printf("before a=%d,b=%d\n",*a,*b)
	*a,*b = *b,*a
	fmt.Printf("after a=%d,b=%d\n",*a,*b)
}

func main()  {
	var a int = 100
	fmt.Printf("a=%d,addr(a)=%p\n",a,&a)
	modifyInt(&a)
	fmt.Printf("a=%d,addr(a)=%p\n",a,&a)
	var b int = 99
	swich(&a,&b)
	fmt.Printf("main a=%d,b=%d\n",a,b)
}
