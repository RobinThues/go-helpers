package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{0}
	if compareIntSlices(a, b) == false {
		fmt.Println("slices correctly recognized as un-equal")
	}
}

