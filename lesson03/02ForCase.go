package main

import "fmt"

func testFor1() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("num %d\n", i)
	}
	fmt.Printf("final num:%d\n", i)
}

func endFor() {
	var i int
	for i = 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("num %d\n", i)
	}
	fmt.Printf("final num %d\n", i)
}

func continueFor() {
	var i int
	for i = 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("num %d is odd\n", i)
	}
	fmt.Printf("final %d\n", i)
}

func testFor2() {
	var i int = 1
	for i <= 10 {
		fmt.Printf("num %d\n", i)
		i += 2
	}
	fmt.Printf("final num %d\n", i)
}

func testFor3() {
	var i int
	for i <= 10 {
		fmt.Printf("num %d\n", i)
		i++
	}
}

func testFor4()  {
	for no,i := 10,1;i <=10 && no <= 19;i,no = i+1,no+1 {
		fmt.Printf("%d * %d = %d\n",no,i,no*i)
	}
}

func deadFor()  {
	for {
		fmt.Printf("hello\n")
	}
}
func main() {
	//testFor1()
	//endFor()
	//continueFor()
	//testFor2()
	//testFor3()
	//testFor4()
	deadFor()
}
