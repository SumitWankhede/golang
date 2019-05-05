package main

import "fmt"

func main() {
	var age = 29
	fmt.Println("my age is ", age)
	age = 26
	fmt.Println("my age is ", age)
	age = age + 2
	fmt.Println("my age is ", age)

	var height, width int = 23, 45
	fmt.Println("Height is ", height, " and width is ", width)

	var ht, wt = 13, 16
	fmt.Println("new height is", ht, "new width is", wt)

	var (
		name     = "Sumit"
		myage    = 29
		lastname = "Wankhede"
	)
	fmt.Println("my name is", name, lastname, "and my age is", myage)

	ini, mini := "ini", 55
	fmt.Println("ini", ini, "mini", mini)
}
