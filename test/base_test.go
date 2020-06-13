package test

import (
	"functional/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	go app.Startapp()
	os.Exit(m.Run())
}