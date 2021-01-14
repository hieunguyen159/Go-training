package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainPanicGoDotEnv(t *testing.T) {
	os.Rename("./.env", "./.env.test")
	assert.Panics(t, main, 0)
	os.Rename("./.env.test", "./.env")
}
func TestMainPanicInitDB(t *testing.T) {
	os.Setenv("PORT", "8080")
	assert.Panics(t, main, 0)
}
