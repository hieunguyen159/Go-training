package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectCubesSuccess(t *testing.T) {
	err := godotenv.Load()
	assert.Equal(t, nil, err)
	var connector = ConnectCubes()
	assert.NotEqual(t, nil, connector)
}
