package giphy

import (
	"github.com/sanzaru/go-giphy"
)

// Init giphy api
func Init(apiKey string) func(tag string) (string, error) {
	giphy := libgiphy.NewGiphy(apiKey)

	// GetRandom image by a given tag
	return func(tag string) (string, error) {
		dataRandom, err := giphy.GetRandom(tag)
		if err != nil {
			return "", err
		}
		return dataRandom.Data.Image_original_url, nil
	}

}
