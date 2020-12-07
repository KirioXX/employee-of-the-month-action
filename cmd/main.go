package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	whoToGreat := argsWithoutProg[0]

	fmt.Printf("Hello %s", whoToGreat)
}
