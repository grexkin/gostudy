package main

import "fmt"

func insert_sort(a [8]int) [8]int {
	for i := 1;i<len(a);i++ {
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

func select_sort(a [8]int) [8]int {
	for i:=0;i<len(a);i++ {
		for j:=i+1;j<len(a);j++ {
			if a[j] < a[i] {
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	return a
}

func buble_sort(a [8]int) [8]int {
	for i := 0;i<len(a);i++ {
		for j:=0;j<len(a)-i-1;j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	return a
}

func main() {
	var i [8]int = [8]int{8,9,23,12,44,1,22,32}
	//插入排序
	j := insert_sort(i)
	fmt.Println("原始数组:         ",i)
	fmt.Println("插入排序之后的值：",j)

	k := select_sort(i)
	fmt.Println("选择排序之后的值：",k)

	l := buble_sort(i)
	fmt.Println("冒泡排序之后的值：",l)

}