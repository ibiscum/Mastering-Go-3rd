// nolint
package main

import (
	"fmt"
)

func S2() {
	return
	fmt.Println("Hello!") //nolint
}

func S1() {
	fmt.Println("In S1()")
	return

	fmt.Println("Leaving S1()") //nolint
}

func main() {
	S1()
}
