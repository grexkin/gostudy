package main

import "fmt"

func justfyPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2;i<n;i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func primeDemo()  {
	//打印100以内的质数
	for i := 2;i < 100;i++ {
		if justfyPrime(i) == true {
			fmt.Printf("[%d] is prime\n" ,i)
		}
	}
}

func isNarcissisticNum(n int) bool {
	digits := n % 10   //取出个位数
	decade := (n/10) % 10 //取出十位数
	hundreds := (n/100) % 10 //取出百位数
	//fmt.Printf("n=%d,digits=%d,decade=%d,hundreds=%d\n",n,digits,decade,hundreds)
	sum := digits*digits*digits+decade*decade*decade+hundreds*hundreds*hundreds
	if n == sum {
		return true
	}
	return false
}

func narcissisticNum() {
	for i:=100;i<1000;i++ {
		if isNarcissisticNum(i) {
			fmt.Printf("%d is narcissisticNum\n",i)
		}
	}
}

func calcStrings(str string)(charCount,numCount,spaceCount,otherCounts int)  {
	//多返回值
	strByte := []rune(str)
	for i := 0;i < len(strByte);i ++ {
		if strByte[i] >= 'a' && strByte[i] <= 'z' || strByte[i] >= 'A' && strByte[i] <= 'Z' {
			charCount++
			continue
		}

		if strByte[i] >= '0' && strByte[i] <= '9' {
			numCount ++
			continue
		}

		if strByte[i] == ' ' {
			spaceCount ++
			continue
		}

		otherCounts ++
	}
	return
}

func main() {
	//primeDemo()
	//narcissisticNum()
	charCount,numCount,spaceCount,otherCount := calcStrings("i am 中国 123455我是哈！！")
	fmt.Printf("charCount:%d\nnumCount:%d\nspaceCount:%d\notherCount:%d\n",charCount,numCount,
		spaceCount,otherCount)
}
