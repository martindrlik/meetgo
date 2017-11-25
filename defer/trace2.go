package main

import (
	"time"
	"log"
)

// START OMIT
func main() {
	defer trace("main")()
	log.Println("main does stuff...")
	time.Sleep(2 * time.Second)
}

func trace(s string) func() {
	log.Println("entering", s)
	start := time.Now()
	return func() {
		elapsed := time.Now().Sub(start)
		log.Printf("leaving %s (took: %v)\n", s, elapsed)
	}
}
// END OMIT