package fortune

import (
	"encoding/json"
	"fmt"
	"log"
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

func TestJsonFormat(t *testing.T) {
	var fortunes Fortunes
	err := json.Unmarshal([]byte(mockingJson()), &fortunes)
	if err != nil {
		log.Fatal(err)
	}

	if len(fortunes.Fortunes) != 2 {
		t.Errorf("Size of Json read is not equal to true size")
	}

	if len(fortunes.Fortunes) != 2 {
		t.Errorf("Size of Json read is not equal to true size")
	}

}

func TestFJsonFile(t *testing.T) {
	quotes := openJson()

	fmt.Println(len(quotes.Fortunes))

}

func TestSelectFortune(t *testing.T) {
	quotes := openJson()
	animal, quote := selectFortune(*quotes)
	fmt.Printf("\nanimal: %v, quote: %v", animal, quote)
}

func TestPickRandom(t *testing.T) {
	x := 0
	for x < 10000000 {
		testInt := pickRandom(11)
		if testInt > 10 {
			t.Errorf("testInt bigger then input")
		}

		x += 1
	}

}
