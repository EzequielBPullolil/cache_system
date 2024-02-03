package cacheManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var uuid = "unique-id"

func init() {
	os.Clearenv()
	os.Setenv("cache_file", "test_cache")

	os.Remove(os.Getenv("cache_file"))
	WriteCache(uuid, "{\"hola\": \"dato\"}")
}

func TestFetchNonRegisteredDataReturnEmptyString(t *testing.T) {
	setupTest()

	data, _ := FetchCache("nonRegisteredUUID")

	assert.Empty(t, data)
}

func TestFetchRegisteredDataReturnAstringWithData(t *testing.T) {
	setupTest()

	data, _ := FetchCache(uuid)

	assert.Contains(t, data, "hola")
	assert.Contains(t, data, "dato")
}
