package fortune

import (
	"encoding/json"
	"fmt"
	"os"
)

type Fortunes struct {
	Fortunes []Fortune `json:"fortunes"`
}

type Fortune struct {
	Animal string   `json:"animal"`
	Quotes []string `json:"quotes"`
}

func openJson() *Fortunes {
	file, err := os.Open("fortunes.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var fortunes Fortunes
	err = json.NewDecoder(file).Decode(&fortunes)
	if err != nil {
		panic(err)
	}

	return &fortunes
}

func FortuneTeller() {
	quotes := openJson()

	for _, quote := range quotes.Fortunes {
		fmt.Printf("\nQuotes from %s:\n", quote.Animal)
		for _, q := range quote.Quotes {
			fmt.Printf("\t%s\n", q)
		}
	}
}
