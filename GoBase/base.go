package main //以package main作为程序的入口点

import "fmt" //导入fmt包

//main函数是入口函数
func main() {
	// vardemo()
	// f1()
	// f2(10)
	fmt.Println(control(8))
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
	//1
	if a > 10 {
		fmt.Println("if a> 10")
	} else {
		fmt.Println("else a> 10")
	}
	//2
	if e := 5; e < a {
		fmt.Println("e := 5; e < a")
	}
	//3
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

	//4
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//5
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		fmt.Println("list[", k, "] = ", v)
	}
	//6
	for pos, char := range "aΦx" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}
	//7
	arrs := [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	for k, v := range arrs {
		for k1, v1 := range v {
			fmt.Println("arrs[", k, ",", k1, "] = ", v1)
		}
	}
	//8
	s0 := []int{0, 0}
	s1 := append(s0, 2, 3, 4, 5)
	for k1, v1 := range s1 {
		fmt.Println("arrs[", k1, "] = ", v1)
	}
	fmt.Println("s1长度：", len(s1))

	//9
	monthdays := map[string]int{"Jan": 31, "Feb": 28, "Mar": 31, "Apr": 30, "May": 31, "Jun": 30, "Jul": 31, "Aug": 31, "Sep": 30, "Oct": 31, "Nov": 30, "Dec": 31}
	delete(monthdays, "Jan")
	for key, day := range monthdays {
		fmt.Println(key, "-", day)
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


close 用于 channel 通讯。使用它来关闭 channel，参阅第 6 章了解更多。
delete 用于在 map 中删除实例。
len和cap 可用于不同的类型，len用于返回字符串、slice和数组的长度。参阅“array、slices 和 map” 小节了解更多关于 slice、数组和函数 cap 的详细信息。
new用于各种类型的内存分配。参阅 “用 new 分配内存” 的第 55 页。
make用于内建类型（map、slice 和 channel）的内存分配。参阅 “用 make 分配内存” 的第 55 页。
copy用于复制 slice。参阅本章的 “slice”。
append 用于追加 slice。参阅本章的 “slice”。
panic和recover 用于 异常 处理机制。参阅 “恐慌（Panic）和恢复（Recover）” 的第 32 页了解更 多信息。
print和println 是底层打印函数，可以在不引入 fmt 包的情况下使用。它们主要用于调试。
complex 、 real和imag 全部用于处理 复数。有了之前给的简单的例子，不用再进一步讨论复数了。

*/
