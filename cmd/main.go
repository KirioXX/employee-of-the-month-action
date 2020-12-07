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
	title := argsWithoutProg[2]

	giphy := libgiphy.NewGiphy(apiKey)

	dataRandom, err := giphy.GetRandom(tag)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Image: %s\n", dataRandom.Data.Image_original_url)
}
