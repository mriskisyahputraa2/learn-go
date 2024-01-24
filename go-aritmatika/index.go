package main

import "fmt"

func main() {
	var (
		a = 100 + 20
		b = 200 - 21
		c = a * 120
		d = b / 20
		e = 10 % 4
	)

	f := 10
	f++

	g := 10
	g--

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
