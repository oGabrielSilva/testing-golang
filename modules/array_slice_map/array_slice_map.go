package main

import "fmt"

func main() {
	// Array
	println("Array:")
	var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arr)

	arr2 := [2]string{"Hello world"}
	fmt.Println(arr2)

	// Slices
	println("\nSlices:")

	var sl = make([]int, 0)
	fmt.Println(sl)

	sl2 := []string{"Hello", "world"}
	fmt.Println(sl2)

	sl2 = append(sl2, "From", "Slice")
	fmt.Println(sl2)

	// Map
	println("\nMaps:")

	var m map[string]bool = map[string]bool{"test": true}
	m["isLoggedIn"] = false

	fmt.Println(m)

	var withMake = make(map[string]int, 0)

	withMake["count"] = 0
	fmt.Println(withMake)
	fmt.Println(fmt.Sprintf("Count = %d", withMake["count"]))

	withMake["test"] = 12
	for key, val := range withMake {
		fmt.Println("key", key, "value", val)
	}

	fmt.Println(len(withMake))
}
