package main

import "fmt"

func appendSlice()  {
	a := []int{1,2,3,4}
	b := []int{5,6,7}
	a = append(a,b...)
	fmt.Printf("b=%v\n",a)
}

func sumSlice(a []int) int {
	sum := 0
	for _,val := range a{
		sum += val
	}
	return sum
}

func testSumSlice()  {
	var a [10]int = [10]int{1,2,3,4,5,6,7,8}
	sum := sumSlice(a[:])
	fmt.Printf("sum=%d\n",sum)

	var b [3]int = [3]int{10,20,30}
	fmt.Println("sum(b)=",sumSlice(b[:]))
}

func modifySlice(a []int)  {
	a[0] = 1000
}

func testModifySlice()  {
	var a[3]int = [3]int{1,2,3}
	modifySlice(a[:])
	fmt.Println(a)
}

func copySlice()  {
	a := []string{"a","b","c","d"}
	b := []string{"e","f","g"}
	copy(a,b)
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	c := []int{1,2}
	d := []int{3,4,5,6}
	copy(c,d)
	fmt.Println(c)
	fmt.Println(d)

}
func main() {
	//appendSlice()
	//testSumSlice()
	//testModifySlice()
	copySlice()
}
