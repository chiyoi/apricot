package logs

import (
	"os"
	"testing"
)

func TestSetOutput(t *testing.T) {
	Info("nyan")
	f, err := os.Create("nyan.txt")
	if err != nil {
		t.Fatal(err)
	}
	SetOutput(f)
	Info("nyan")
	Info("nyan")
}

func TestPrependPrefix(t *testing.T) {
	Info("nyan")
	Prefix("[neko] ")
	Info("nyan")
}
