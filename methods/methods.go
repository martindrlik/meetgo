package main

import (
	"fmt"
	"time"
)

type BornNamed struct {
	Born int
	Name string
}

func (bn BornNamed) Age() int {
	return time.Now().Year() - bn.Born
}

var meganFox = BornNamed{1986, "Megan Fox"}

func main() {
	fmt.Printf("Megan is %d years old.", meganFox.Age())
}
