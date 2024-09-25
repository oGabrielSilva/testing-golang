package main

import "fmt"

func main() {
	// Valores de variáveis vazias

	var i int
	fmt.Printf("Valor vazio de um %T -> %v\n", i, i)

	var f float32
	fmt.Printf("Valor vazio de um %T -> %v\n", f, f)

	var f64 float64
	fmt.Printf("Valor vazio de um %T -> %v\n", f64, f64)

	var b bool
	fmt.Printf("Valor vazio de um %T -> %v\n", b, b)

	var s string
	fmt.Printf("Valor vazio de um %T -> \"%v\"\n", s, s)
}
