package main

import (
	"flag"
	"log"
)

func main() {
	var up, down bool
	flag.BoolVar(&up, "up", false, "bring migrations up")
	flag.BoolVar(&down, "down", false, "bring migrations down")
	flag.Parse()

	if (up && down) || (!up && !down) {
		log.Fatal("please set either -up or -down flags")
	}

	switch {
	case up:
		log.Println("migrate up")
	case down:
		log.Println("migrate down")
	}
}
