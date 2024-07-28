// config.go
package config

import (
	"os"
)

var (
	DukcapilAPIKey = os.Getenv("DUKCAPIL_API_KEY")
	DukcapilAPIURL = "https://api.dukcapil.kemendagri.go.id"
)
