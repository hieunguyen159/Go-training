package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDBSuccess(t *testing.T) {
	// err := godotenv.Load()
	// assert.Equal(t, nil, err)
	cubesColDB := NewCollectionDB("Cube")
	assert.NotEqual(t, nil, cubesColDB.c)
}
