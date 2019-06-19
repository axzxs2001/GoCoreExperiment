package main //以package main作为程序的入口点

import "fmt" //导入fmt包

//main函数是入口函数
func main() {
	// 数组
	var arr []string = []string{"a", "b", "c", "d"}
	fmt.Println("数组长度是：", len(arr))
	// 切片
	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice长度是：", len(slice))
	ss := append(slice, 11, 12)
	fmt.Println("slice长度是：", len(ss))
	// 通道
	var ch = make(chan int, 10)
	ch <- 1
	ch <- 12
	ch <- 23
	ch <- 34
	fmt.Println("ch：", <-ch)
	fmt.Println("ch：", <-ch)
	fmt.Println("chan长度是：", len(ch))
	// 字典
	var mp map[string]bool = map[string]bool{
		"0": false,
		"1": true,
	}
	fmt.Println("map长度是：", len(mp))
	// 字符串
	var str string = "道德经"
	fmt.Println("string类型长度是：", len(str))
	fmt.Println("rune类型长度是：", len([]rune(str)))
}
