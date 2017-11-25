package main

import (
	"time"
	"log"
)

// START OMIT
func main() {
	defer un(trace("main"))
	log.Println("main does stuff...")
	time.Sleep(2 * time.Second)
}
// END OMIT

func un(s string, start time.Time) {
	log.Printf("leaving %s (took %v)\n", s, time.Now().Sub(start))
}

func trace(s string) (string, time.Time) {
	log.Println("entering", s)
	return s, time.Now()
}