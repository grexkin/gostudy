package main

import "fmt"

//冒泡排序
func buble_sort(a [8]int)[8]int {
	for i:=0;i<len(a);i++ {
		for j:=0;j<len(a)-i-1;j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	return a
}

//插入排序
func insert_sort(a [8]int) [8]int {
	for i:=1;i<len(a);i++ {
		for j:=i;j>0;j-- {
			if a[j] < a[j-1] {
				a[j],a[j-1] = a[j-1],a[j]
			}else {
				break
			}
		}
	}
	return a
}

//选择排序
func select_sort(a [8]int) [8]int{
	for i := 0;i < len(a);i++ {
		for j:=i+1;j < len(a);j++ {
			if a[j] < a[i] {
				a[j],a[i] = a[i],a[j]
			}
		}
	}
	return a
}

func main() {
	a := [8]int{12,34,45,2,43,100,90,33}
	b := buble_sort(a)
	fmt.Println(b)
	c := [8]int{12,34,45,2,43,100,90,33}
	d := insert_sort(c)
	fmt.Println(d)
	e := [8]int{12,34,45,2,43,100,90,33}
	f := select_sort(e)
	fmt.Println(f)
}
