package embedding

import _ "embed"

//go:embed fortunes.json
var JsonFile []byte
