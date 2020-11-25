package main

import (
	"os"
	"testing"
)

func setup() {

}

func shutdown() {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
