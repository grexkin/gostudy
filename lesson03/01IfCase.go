package main

import "fmt"

func iftest1() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num %d is even\n", num)
	} else {
		fmt.Printf("num %d is odd\n", num)
	}
}

func iftest2() {
	num := 10
	if num > 5 && num < 10 {
		fmt.Printf("num %d is >5 and < 10\n", num)
	} else if num >= 10 && num <= 20 {
		fmt.Printf("num %d is > 10 and < 20\n", num)
	} else if num >= 20 && num <= 30 {
		fmt.Printf("num %d is > 20 and < 30\n", num)
	} else {
		fmt.Printf("num %d is > 30\n", num)
	}
}

func iftest3() {
	//num 的生命周期只是在if语句块当中生效，在if之外就会被销毁
	if num := 10; num%2 == 0 {
		fmt.Printf("num %d is even\n", num)
	} else {
		fmt.Printf("num %d is odd\n", num)
	}
}

func getNum() int {
	return 10
}

func iftest4() {
	if num := getNum(); num%2 == 0 {
		fmt.Printf("num %d is even\n", num)
	} else {
		fmt.Printf("num %d is odd\n", num)
	}
}

func main() {
	//iftest1()
	//iftest2()
	//iftest3()
	iftest4()
}
