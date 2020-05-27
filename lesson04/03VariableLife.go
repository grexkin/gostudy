package main

import "fmt"

var a int = 100
func testGlobal()  {
	fmt.Printf("a=%d\n",a)
}

func testLocal()  {
	var a int = 1999
	var b int = 10
	fmt.Printf("a=%d,b=%d\n",a,b)
	if b == 10 {
		var c int = 100
		fmt.Printf("c=%d\n",c)
	}

	if d := 100;d>0 {
		fmt.Printf("d=%d\n",d)
	}else {
		fmt.Printf("d=%d\n",d)
	}

	for i :=0;i<10;i++ {
		fmt.Printf("i=%d\n",i)
	}
}

func main() {
	testGlobal()
	testLocal()
}
