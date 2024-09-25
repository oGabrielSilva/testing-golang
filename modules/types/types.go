package main

import "fmt"

func main() {
	fmt.Printf("Value: %t; Type: %T\n", true, true)
	fmt.Printf("Value: %s; Type: %T\n", "String", "")
	fmt.Printf("Value: %d; Type: %T\n", 11, 11)
	fmt.Printf("Value: %f; Type: %T\n", float32(11.5), float32(11.5))
	fmt.Printf("Value: %f; Type: %T\n", 11.5, 11.5)
}
