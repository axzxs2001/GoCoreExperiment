package main //以package main作为程序的入口点

import "fmt" //导入fmt包

//main函数是入口函数
func main() {
	// vardemo()
	// f1()
	// f2(10)
	fmt.Println(control(2))
}

//变量测试
func vardemo() {
	//定义赋值变量方法一
	var i int = 10
	fmt.Println("i =", i)
	//定义赋值变量方法二
	var f32 float32
	f32 = 1.234
	fmt.Println("f =", f32)
	//定义赋值变量方法三
	b := true
	fmt.Println("b =", b)

	var (
		x int
		y float64
	)
	x = 123
	y = 2.345
	fmt.Println("x =", x, ",y =", y)

	a1, a2 := true, "这是a2的值"
	fmt.Println("a1=", a1, ",a2 =", a2)

	var s = "abc" +
		"\n123"
	fmt.Println("s=", s)
}

func control(a int) int {
	if a > 10 {
		fmt.Println("if a> 10")
	} else {
		fmt.Println("else a> 10")
	}
	switch a {
	case 1:
		fmt.Println("switch case 1")
		break
	case 2:
		fmt.Println("switch case 2")
		break
	default:
		fmt.Println("switch case default")
		break
	}
	return 10
}
func f1() {
	fmt.Println("func f1")
}
func f2(i int) {
	fmt.Printf("func f2(%d)", i)
}

/*
break
default
func
interface
select
case
defer
go
map
struct
chan
else
goto
package
switch
const
fallthrough
if
range
type
continue
for
import
return
var

*/
