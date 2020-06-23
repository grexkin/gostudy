package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func demo1()  {
	var sa = make([]string,5,10)
	for i:=0;i<10;i++ {
		sa=append(sa,fmt.Sprintf("%d",i))
	}
	fmt.Println(sa)
}

func sortArray()  {
	a := [10]int{19,123,243,243,323,24,12,45,19,33}
	sort.Ints(a[:])
	fmt.Println(a)
	b := []string{"v","a","i","o","k"}
	sort.Strings(b)
	fmt.Println(b)
	c := []float64{3.12,4.22,90.11,22.11,88}
	sort.Float64s(c)
	fmt.Println(c)
}

func geneChar()  {
	l := 'A'
	for i:=0;i<26;i++ {
		fmt.Printf("%c",l)
		l+=1
	}
	s := 'a'
	for j:=0;j<26;j++ {
		fmt.Printf("%c",s)
		s+=1
	}
}

var (
	numCharset = "0123456789"
	strCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	mixCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	advanceCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()"
)

func main() {
	//demo1()
	//sortArray()
	//geneChar()
	var length int
	var charset string
	
	flag.IntVar(&length,"l",16,"-l the length of passwd")
	flag.StringVar(&charset,"t","mix","-t the charset of passwd(num|char|mix|advance) ")
	flag.Parse()
	
	rand.Seed(time.Now().UnixNano())
	
	var userCharset string     //定义密码字符串

	switch charset {
	case "num":
		userCharset = numCharset
	case "char":
		userCharset = strCharset
	case "mix":
		userCharset = mixCharset
	case "advance":
		userCharset = advanceCharset
	default:
		userCharset = mixCharset
	}

	var password []byte    //定义存密码的切片
	for i:=0;i<length;i++ {
		index := rand.Intn(len(userCharset))
		char := userCharset[index]
		password = append(password,char)
	}

	strPassword := string(password)
	fmt.Printf("生成的密码是%s\n",strPassword)
}
