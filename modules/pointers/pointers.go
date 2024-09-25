package main

import "fmt"

func main() {
	x := 11
	y := x

	println("func main ========")
	y = 22
	fmt.Println(x, y)
	fmt.Println(&x, &y)
	p(x, y)

	withRef()
}

func p(x, y int) {
	println("\nfunc p ========")
	fmt.Println(x, y)
	fmt.Println(&x, &y)
}

func pPoint(x, y *int) {
	println("\nfunc pPoint ========")
	fmt.Println(x, y)
	fmt.Println(*x, *y)
}

func withRef() {
	x := 11
	y := &x

	println("\nfunc withRef ========")
	fmt.Println("X:", x, "Y:", *y)
	fmt.Println(&x, y)

	*y = 12
	fmt.Println("X:", x, "Y:", *y)
	fmt.Println(&x, y)

	pPoint(&x, y)
}
