package main

import "fmt"

var a = 6

func main() {
	p()
	q()
	p()
	r, m := abc()
	fmt.Println("r=", r, ",m=", m)
	//异常
	ff := func() {
		arr := []string{"元素0"}
		arr[1] = "元素1"
		//主动抛异常
		panic("这里抛出一个异常")
	}
	b := throwsPanic(ff)
	fmt.Println("b=", b)
}

func abc() (result bool, message string) {
	result = false
	message = "返回错误"
	return
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if exc := recover(); exc != nil {
			fmt.Println(exc)
			b = true
		}
	}()
	f()
	return
}

func p() {
	fmt.Println(a)
}

func q() {
	a = 5
	fmt.Println(a)
}
