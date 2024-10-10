package fortune

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func mockingJson() string {
	jsonData := `{
    "fortunes": [
        {
            "animal": "turtle",
            "quotes": [
                "There is strength in taking it slow... and I’m really strong!",
                "Slow and steady gets you there – with way less stress.",
                "Haste makes waste, but really... let waste win."
            ]
        },
        {
            "animal": "wolf",
            "quotes": [
                "A lone wolf still howls at the moon, even when no one is listening.",
                "Strength doesn’t come from numbers; it comes from the fight within.",
                "The pack protects, but the lone wolf thrives in silence."
            ]
        }
    ]
}`

	return jsonData
}

func mockingOpenJson() *Fortunes {
	file, err := os.Open("../../fortunes.json")
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

func TestOpenJson(t *testing.T) {
	file, err := os.Open("../../fortunes.json")
	if err != nil {
		t.Errorf("Failed to open Json: %v", err)
	}

	defer file.Close()

	var quotes Fortunes
	err = json.NewDecoder(file).Decode(&quotes)
	if err != nil {
		t.Errorf("Failed to process Json: %v", err)
	}
}

func TestJsonFormat(t *testing.T) {
	var fortunes Fortunes
	err := json.Unmarshal([]byte(mockingJson()), &fortunes)
	if err != nil {
		log.Fatal(err)
	}
}

func TestFJsonFile(t *testing.T) {
	quotes := mockingOpenJson()

	for _, quote := range quotes.Fortunes {
		fmt.Printf("\nQuotes from %s:\n", quote.Animal)
		for _, q := range quote.Quotes {
			fmt.Printf("\t%s\n", q)
		}
	}

}
