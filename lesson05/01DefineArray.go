package main

import "fmt"

func defineArray()  {
	var a [3]int
	a = [3]int{1,2,3}
	fmt.Println(a)
	a[0] = 100
	fmt.Println(a)

	var b [4]int
	b[0] = 3
	b[1] = 4
	for i,j := range b {
		fmt.Println(i,j)
	}

	c := [...]int{1,2,3,4,5,6}
	fmt.Println(c,len(c))

	//对下标元素初始化
	d := [5]int{3:100,4:200}
	fmt.Println(d)
}

func defineArray2(){
	var a [2]int = [2]int{3,4}
	var b [2]int = [2]int{1,2}
	b = a
	fmt.Println(b)
}

func loopArray()  {
	var a [5]int = [5]int{2,3,4,3:100,4:200}
	for i:=0;i<len(a);i++ {
		fmt.Println(a[i])
	}
}

func loopArrayRange()  {
	var a [5]int = [5]int{2,3,4,3:100,4:200}
	for index,val := range a {
		fmt.Printf("index:%d,val:%v\n",index,val)
	}
}

func dubleArray()  {
	var a[3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40
	a[2][0] = 50
	a[2][1] = 60
	//fmt.Println(a)
	for i:=0;i<3;i++{
		for j:=0;j<2;j++{
			fmt.Printf("%d\t",a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("======")
	for i,val := range a {
		for j,val2 := range val {
			fmt.Printf("(%d,%d)=%d ",i,j,val2)
		}
		fmt.Println()
	}
}

func dubleArray2()  {
	a := [3][2]string{
		{"aaa","AAA"},
		{"bbb","BBB"},
		{"ccc","CCC"},
	}
	for _,v1 := range a {
		for _,v2 := range v1 {
			fmt.Printf("%s ",v2)
		}
		fmt.Println()
	}
}

func main() {
	defineArray()
	defineArray2()
	loopArray()
	loopArrayRange()
	dubleArray()
	dubleArray2()
}