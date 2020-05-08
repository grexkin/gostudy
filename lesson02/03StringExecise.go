package main

import "fmt"

func reverseEn(s string) string  {
	strByte := []byte(s)
	for i:=0;i<len(s)/2;i++ {
		strByte[i],strByte[len(s)-i-1] = strByte[len(s)-i-1],strByte[i]
	}
	s = string(strByte)
	return s
}

func reverseCh(s string)  {
	strRune := []rune(s)
	for i:=0;i< len(strRune)/2;i++ {
		strRune[i],strRune[len(strRune)-i-1] = strRune[len(strRune)-i-1],strRune[i]
	}
	s = string(strRune)
	fmt.Println(s)
}

func plalindrome(s string)string {
	//回文判断程序
	strRune := []rune(s)
	for i:=0;i<len(strRune);i++ {
		strRune[i],strRune[len(strRune)-i-1] = strRune[len(strRune)-i-1],strRune[i]
	}
	s1 := string(strRune)
	fmt.Println(s)
	fmt.Println(s1)
	if s1 == s {
		return "plalindrome"
	} else {
		return "is not plalindrome"
	}
}

func main () {
	EnStr := "abcdefghi jkl"
	fmt.Println(reverseEn(EnStr))
	s := "鸡哥like大宝剑"
	reverseCh(s)
	s1 := "大和尚尚和大1sss"
	fmt.Println(plalindrome(s1))
}
