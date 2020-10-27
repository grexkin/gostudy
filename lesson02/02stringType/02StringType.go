package main

import "fmt"

func loopString() {
	str := "hello abc"
	fmt.Printf("str[0]=%c,len(str)=%d\n",str[0],len(str))
	// for loop
	for i := 0;i < len(str);i++ {
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}
	//for range
	for index,val := range str {
		fmt.Printf("str[%d]=%c\n",index,val)
	}
}

func modifyString() {
	//纯字母类型的修改
	str := "hello"
	strByte := []byte(str)    //转换成字符切片
	strByte[0] = 'c'          //修改第一个元素
	str = string(strByte)    //转换成字符串
	fmt.Printf("str=%s\n",str)
	//带中文的字符串修改
	str_ch := "hello大帅哥"
	str_chByte := []rune(str_ch)
	str_chByteb := []rune(str_ch)
	fmt.Printf("len(str_chByte)=%d,len(str_chByteb)=%d\n",len(str_chByte),len(str_chByteb))
	str_chByte[5] = 'z'
	str_chByte[6] = '大'
	str_ch = string(str_chByte)
	fmt.Printf("str_ch=%s\n",str_ch)
}
func main() {
	//loopString()
	modifyString()
}