package main

import "fmt"

func main() {
	n := 42
	p := &n
	fmt.Printf("%T\n", p)

	*p = 54
	fmt.Println(n)
}
