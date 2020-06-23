package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length int
	charset string
	sourceChar string
)

const (
	NumStr  = "0123456789"
	CharStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MixStr  =  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AdvanceStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!~#$%^&*()_+=,."
)

func parserArgs()  {
	flag.IntVar(&length,"l",16,"-l 生成密码的长度")
	flag.StringVar(&charset,"t","mix",`
				-t 指定密码复杂度，
                num: 只有数字复杂度
				char: 字母大小写复杂度
				mix: 字母数字混合复杂度
				advande: 字母数字特殊符号
    `)
	flag.Parse()
}

//生成字符串
func genStr()  {
	l := 'a'
	for i:=0;i<26;i++ {
		fmt.Printf("%c",l)
		l+=1
	}
	m := 'A'
	for j:=0;j<26;j++ {
		fmt.Printf("%c",m)
		m+=1
	}
}

func generaPasswd() string {
	//初始化随机种子
	rand.Seed(time.Now().UnixNano())
	switch charset {
	case "num":
		sourceChar = NumStr
	case "char":
		sourceChar = CharStr
	case "mix":
		sourceChar = MixStr
	case "advance":
		sourceChar = AdvanceStr
	default:
		sourceChar = MixStr
	}

	var password []byte = make([]byte,0,length)
	for i :=0 ;i <length;i ++ {
		index := rand.Intn(len(sourceChar))
		char := sourceChar[index]
		password = append(password,char)
	}

	return string(password)
}

func main() {
	parserArgs()
	fmt.Printf("length:%d,charset:%s\n",length,charset)
	passwd := generaPasswd()
	fmt.Printf("生成的密码是： %s\n",passwd)
}