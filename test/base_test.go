package test

import (
	"github.com/namantalaycha/functional/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	go app.Startapp()
	os.Exit(m.Run())
}