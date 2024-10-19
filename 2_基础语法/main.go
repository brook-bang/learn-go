package main

import "fmt"

func main() {
	fmt.Println("Google" + "Runoob")

	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	fmt.Printf(url, stockcode, enddate)

}
