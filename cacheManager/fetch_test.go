package cacheManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var uuid = "unique-id"

func init() {
	os.Clearenv()
	os.Remove(".cache_file")
}

func TestFetchNonRegisteredDataReturnEmptyString(t *testing.T) {
	setupTest()
	data, _ := FetchCache("nonRegisteredUUID")
	assert.Empty(t, data)
}

func TestFetchRegistered(t *testing.T) {
	setupTest()
	data, _ := FetchCache(uuid)
	assert.Empty(t, data)
	WriteCache(uuid, "hola")

	data, _ = FetchCache(uuid)

	assert.Equal(t, data, "data=hola")
}
