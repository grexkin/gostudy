package _3stringExe

import "fmt"

func reverseEn(s string) string  {
	strByte := []byte(s)
	for i := 0;i < len(s)/2; i++ {
		strByte[i],strByte[len(s)-i-1] = strByte[len(s)-i-1],strByte[i]
	}
	s = string(strByte)
	return s
}

func reverseCh(s string) string {
	strByte := []rune(s)
	for i:=0; i < len(strByte)/2; i++ {
		strByte[i],strByte[len(strByte)-i-1] = strByte[len(strByte)-i-1],strByte[i]
	}
	s = string(strByte)
	return s
}

func plalindrome()  {
	//回文判断程序
	s := "上海自来水好来自海上"
	strRune := []rune(s)
	for i:=0;i<len(strRune);i++ {
		strRune[i],strRune[len(strRune)-i-1] = strRune[len(strRune)-i-1],strRune[i]
	}
	s1 := string(strRune)
	fmt.Println(s)
	fmt.Println(s1)
	if s1 == s {
		fmt.Printf("%s is plalindrome\n",s)
	} else {
		fmt.Printf("%s is not plalindrome\n",s)
	}
}

func main () {
	EnStr := "abcdefghi jkl"
	fmt.Println(reverseEn(EnStr))
	fmt.Println(reverseCh("大河向东flowing！"))
	s := "鸡哥like大宝剑"
	reverseCh(s)

	plalindrome()
}
