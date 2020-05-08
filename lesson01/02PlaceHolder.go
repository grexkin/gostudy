package main

import "fmt"

func placeHolder()  {
	var a int = 19
	fmt.Printf("a=%d\n",a)
	b := 01011
	fmt.Printf("b=%T\n",b)
	var (
		c = "hello"
		d = 3.1415926
	)
	fmt.Printf("c=%s,d=%.2f\n",c,d)
	e := 1234
	fmt.Printf("e=%b,e=%x,e=%d\n",e,e,e)
}


func main() {
	placeHolder()
}
