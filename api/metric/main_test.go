package metric

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	code := m.Run()
	os.Exit(code)
}
