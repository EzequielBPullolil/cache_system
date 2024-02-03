package cacheManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTest() {
	os.Clearenv()
	os.Setenv("cache_file", "test_cache")

	os.Remove(os.Getenv("cache_file"))
}

func TestWriteOperationModifyCacheFile(t *testing.T) {
	setupTest()
	data, _ := os.ReadFile(os.Getenv("cache_file"))

	assert.Empty(t, data)

	WriteCache("test_uuid", "fake_data")
	data, _ = os.ReadFile(os.Getenv("cache_file"))

	assert.Contains(t, string(data), "test_uuid")
	assert.Contains(t, string(data), "fake_data")
}

func TestWriteOperationPersistOldData(t *testing.T) {
	setupTest()
	data, _ := os.ReadFile(os.Getenv("cache_file"))
	assert.Empty(t, data)

	fake_data := []string{"hola", "esto", "es", "info", "falsa"}
	for i := 0; i < len(fake_data); i++ {
		WriteCache("id", fake_data[i])
	}
	data, _ = os.ReadFile(os.Getenv("cache_file"))
	for i := 0; i < len(fake_data); i++ {
		assert.Contains(t, string(data), fake_data[i])
	}
}
