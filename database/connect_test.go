package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionSuccess(t *testing.T) {
	var cubesConnector = ConnectCubes()
	assert.NotEqual(t, nil, cubesConnector)
	var emailsConnector = ConnectEmails()
	assert.NotEqual(t, nil, emailsConnector)
	var usersConnector = ConnectUsers()
	assert.NotEqual(t, nil, usersConnector)
}
