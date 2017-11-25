package main

import "fmt"

type BornNamed struct {
	Born int
	Name string
}

var meganFox = BornNamed{Born: 1986, Name: "Megan Fox"}

func main() {
	fmt.Println(meganFox)
}
