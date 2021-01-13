package db

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInitDBSuccess(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Equal(t, nil, err)
	// cubesColDB := NewCollectionDB("Cube")
	// assert.NotEqual(t, nil, &cubesColDB.c)
}
