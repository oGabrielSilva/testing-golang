package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(withParam("world"))
	fmt.Println(returnsNamed("Hello, world!"))
	fmt.Println(returnsNamed("Hello, world."))
	fmt.Println(withFunc(12)())

	oneType("A", "BBB", "CCC", 1, 2, 3)
}

func returnsNamed(payload string) (containsDot bool, containsX bool) {
	containsDot = strings.Contains(payload, ".")
	containsX = strings.Contains(payload, "x")
	return
}

func withParam(p string) string {
	return "Hello, " + p
}

func withFunc(data int) func() int {
	return func() int {
		return data * 2
	}
}

func oneType(a, b, c string, d, e, j int) {
	fmt.Println("Todas abaixo possuem o mesmo tipo:")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c + "\n")

	fmt.Println("Todas abaixo possuem o mesmo tipo:")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(fmt.Sprint(j) + "\n")
}
