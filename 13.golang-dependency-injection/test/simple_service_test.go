package test

import (
	"fmt"
	"golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)

	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestDatabaseRepository(t *testing.T) {
	databaseRepository := simple.InitializedDatabaseRepository()

	fmt.Println(*databaseRepository)
	assert.NotNil(t, databaseRepository)
}

func TestConnection(t *testing.T) {
	connect, cleanup := simple.InitializedConnection("Test")

	fmt.Println(connect)

	cleanup()
}