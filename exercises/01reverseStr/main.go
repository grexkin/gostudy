package main

import "fmt"

func reverse_str(str string) string  {
	s_origin := str
	str_rune := []rune(str)
	for i:=0;i<len(str_rune)/2;i++ {
		str_rune[i],str_rune[len(str_rune)-i-1] = str_rune[len(str_rune)-i-1],str_rune[i]
	}
	fmt.Println(s_origin)
	s_reversed := string(str_rune)
	if s_reversed == s_origin {
		fmt.Printf("%s is 回文",str)
	}else {
		fmt.Printf("%s 不是回文",str)
	}
	return s_reversed
}

func main() {
	s := "abcdef中国"
	s_r := reverse_str(s)
	fmt.Println(s_r)
}
