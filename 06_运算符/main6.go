package main

import "fmt"

func main() {

	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d
	fmt.Printf("e = (a + b) * c / d的值为：%d\n", e)

	e = ((a + b) * c) / d
	fmt.Printf("e = ((a + b) * c) / d的值为：%d\n", e)

	e = (a + b) * (c / d)
	fmt.Printf("e = (a + b) * (c / d)的值为：%d\n", e)

	e = a + (b*c)/d
	fmt.Printf("e = a + (b * c) / d的值为：%d\n", e)

}
