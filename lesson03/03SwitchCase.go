package main

import "fmt"

func beforeSwitch()  {
	var a int = 2
	if a == 1 {
		fmt.Printf("a=1\n")
	} else if  a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 3 {
		fmt.Printf("a=3\n")
	} else if a == 4 {
		fmt.Printf("a=4\n")
	} else {
		fmt.Printf("other")
	}
}

func switchCase1()  {
	var a int = 3
	switch a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
		fallthrough
	case 4:
		fmt.Printf("a=4\n")
	default:
		fmt.Printf("others\n")
	}
}

func switchCase2()  {
	//等同于switchCase1
	switch a :=2;a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
		fallthrough
	case 4:
		fmt.Printf("a=4\n")
	default:
		fmt.Printf("others\n")
	}
}

func switchCase3() {
	var num int = 21
	switch {
	case num > 1 && num < 10:
		fmt.Printf("num %d > 1 && < 10",num)
	case num > 10 && num < 20:
		fmt.Printf("num %d > 10 && < 20",num)
	case num > 20 && num < 30:
		fmt.Printf("num %d > 20 && < 30",num)
		fallthrough
	default:
		//fallthrough 不能穿透到default
		fmt.Printf(" others\n")
	}
}

func switchCase4()  {
	letter := "i"
	switch letter {
	case "a","e","i","o","u":
		fmt.Printf("vowel\n")
	default:
		fmt.Printf("not a vowel")
	}
}
func main() {
	//beforeSwitch()
	//switchCase1()
	//switchCase2()
	//switchCase3()
	switchCase4()

}
