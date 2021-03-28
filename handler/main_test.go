package handler

import (
	"os"
	"testing"

	"github.com/shyam81992/Weather-Monster/config"
)

func TestMain(t *testing.M) {
	exitVal := t.Run()

	if os.Getenv("integration_testing") == "true" {
		config.LoadConfig()
	}
	os.Exit(exitVal)
}
