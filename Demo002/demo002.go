package main

import "fmt"

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5, 6}

	b := make([]int, 2, 4)

	c := a[0:3]

	fmt.Printf("%d\r\n", len(b))
	fmt.Printf("%d\r\n", cap(b))
	fmt.Println(b)
	fmt.Println(c)
	c = append(c, 100)
	fmt.Println(c)
	fmt.Printf("%d\r\n", len(c))
	fmt.Printf("%d\r\n", cap(c))
}
