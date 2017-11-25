package main

import (
	"fmt"
	"time"
)

type Born int

func (b Born) Age() int {
	return time.Now().Year() - int(b)
}

type BornNamed struct {
	Born // HL
	Name string
}

var meganFox = BornNamed{1986, "Megan Fox"}

func main() {
	fmt.Println(meganFox.Age())
}
