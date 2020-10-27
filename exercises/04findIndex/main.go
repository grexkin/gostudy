package main

import (
	"fmt"
	"math/rand"
	"time"
)

//所有元素之和
func calc_array_sum(){
	var arr [10]int
	for i:=0;i<len(arr);i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	sum := 0
	for j := 0;j<len(arr);j++ {
		sum += arr[j]
	}
	fmt.Println(sum)
}

//元素之和等于12 的下标

func index_return(){
	var arr [10]int
	for i :=0;i<len(arr);i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	for j:=0;j<len(arr);j++ {
		for k:=j+1;k<len(arr);k++ {
			if (arr[j]+arr[k] == 12) {
				fmt.Printf("下标是j=%d,k=%d\n",j,k)
			}
		}
	}
}

func date_time()  {
	now := time.Now()
	fmt.Println(now)
}

func main() {
	calc_array_sum()
	index_return()
	date_time()
}
