package main

import (
	"fmt"
	"strings"
)

func variablesType()  {
	var a int = 100
	b,c,d,e := 3.1415926,false,"hello","中国"
	var f string
	f = "ass"
	fmt.Printf("a=%d,b=%.3f,c=%t,d=%s,e=%s,f=%s\n",a,b,c,d,e,f)
}

func constDemo()  {
	//常量练习
	const a int = 100
	const b string = "吃喝拉撒"
	//ab := a + b
	fmt.Println(b)
	const (
		c = iota
		d
		e
		f = 10
		g
		h = "hello"
		i
		j = iota   //这里不等于0
		k
		l
		m
	)
	const (
		n = iota
		o
		p
	)
	fmt.Println(c,d,e,f,g,h,i,j,k,l,m,n,o,p)
}

func boolType() {
	var a bool = true
	b := false
	fmt.Println("a =",a,"b =",b)
	c := a && b
	d := a || b
	fmt.Printf("c=%t\nd=%t\n",c,d)
}

func intFloatType() {
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f uint
		g uint8
		h uint16
		i uint32
		j uint64
		k float32
		l float64
	)
	fmt.Println(a,b,c,d,e,f,g,h,i,j,k,l)
}

func formatStdout()  {
	//原样输出字符串
	a := `\n\n\t<html>aaa</html>
            佛祖保佑，永无bug
			阿弥陀佛
            010100101
          `
	fmt.Println(a)
	//字符串拼接
	b := "aaa"
	c := "bbb"
	d := b + c
	fmt.Println(d)
	e := fmt.Sprintf("%s-%s",b,c)
	fmt.Println(e)
	//字符串的其他操作
	ips := "192.168.56.11,192.168.56.12,192.168.56.13"
	ipList := strings.Split(ips,",")
	fmt.Println(ipList[1])
	//contain
	ipContain := strings.Contains(ips,"192.168.56.14")
	fmt.Println(ipContain)
	//前缀判断
	url := "http://www.baidu.com"
	if strings.HasPrefix(url,"http") {
		fmt.Println("this is a http web")
	}else {
		fmt.Println("this is not http web")
	}
	if strings.HasSuffix(url,"com") {
		fmt.Println("this is endupwith com")
	}else {
		fmt.Println("this is not endupwith com")
	}

	index := strings.Index(url,"baidu")
	fmt.Println("baidu in position",index)

	var strArry = []string{"192.168.56.11","192.168.56.12"}
	res := strings.Join(strArry,"===")
	fmt.Printf("joindStr=%s\n",res)

}

func main() {
	//variablesType()
	constDemo()
	//boolType()
	//intFloatType()
	formatStdout()
}