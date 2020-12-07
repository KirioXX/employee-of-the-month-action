package main

import (
	"fmt"
	"github.com/sanzaru/go-giphy"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	apiKey := argsWithoutProg[0]
	tag := argsWithoutProg[1]

	giphy := libgiphy.NewGiphy(apiKey)

	dataRandom, err := giphy.GetRandom(tag)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Random data: %+v\n", dataRandom)
}
