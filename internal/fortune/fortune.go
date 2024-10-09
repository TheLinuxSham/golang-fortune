package fortune

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Fortune struct {
	Animal string   `json:"animal"`
	Quotes []string `json:"quotes"`
}

type Content struct {
	Content []Fortune `json:"content"`
}

func FortuneTeller() {
	file, err := os.Open("fortunes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var quotes Content
	err = json.NewDecoder(file).Decode(&quotes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quotes:")
	for _, quote := range quotes.Content {
		fmt.Printf("\nQuotes from %s:\n", quote.Animal)
		for _, q := range quote.Quotes {
			fmt.Printf("\t%s\n", q)
		}
	}
}
