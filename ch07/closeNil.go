//go:build ignore

package main

func main() {
	var c chan string
	close(c)
}
