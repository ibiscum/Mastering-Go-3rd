//go:build ignore

package main

import (
	_ "embed"
	"fmt"
)

//go:embed printSource.go
var src string

func main() {
	fmt.Print(src)
}
