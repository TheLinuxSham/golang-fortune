package fortune

import (
	"encoding/json"
	"math/rand"

	"github.com/TheLinuxSham/golang-fortune/internal/embedding"
)

type Fortunes struct {
	Fortunes []Fortune `json:"fortunes"`
}

type Fortune struct {
	Animal string   `json:"animal"`
	Quotes []string `json:"quotes"`
}

func pickRandom(maxValue int) int {
	return rand.Intn(maxValue)
}

func openJson() *Fortunes {
	// file, err := os.Open(filepath)
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	var fortunes Fortunes
	err := json.Unmarshal(embedding.JsonFile, &fortunes)
	if err != nil {
		panic(err)
	}

	return &fortunes
}

func selectFortune(fortunes Fortunes) (string, string) {
	idFortunes := pickRandom(len(fortunes.Fortunes))
	idFortune := pickRandom(len(fortunes.Fortunes[idFortunes].Quotes))

	animal := fortunes.Fortunes[idFortunes].Animal
	quote := fortunes.Fortunes[idFortunes].Quotes[idFortune]

	return animal, quote
}

func FortuneTeller() (string, string) {
	quotes := openJson()
	animal, quote := selectFortune(*quotes)

	return animal, quote
}
