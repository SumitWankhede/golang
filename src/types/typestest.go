package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false

	fmt.Println("a:", a, "b:", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d:", d)

	var e = 89
	f := 95
	fmt.Println("value of e is", e, "and f is", f)
	fmt.Printf("type of e %T, size of e is %d\n", e, unsafe.Sizeof(e))
	fmt.Printf("type of f is %T, size is f is %d\n", f, unsafe.Sizeof(f))

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)

	k := 10
	var l = float64(k) //this statement will not work without explicit conversion
	fmt.Println("l", l)
}
