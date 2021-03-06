package utils_test

import (
	"crawler/utils"
	"os"
	"testing"
)

func TestMongodbClient(t *testing.T) {
	os.Setenv("MONGODB_USERNAME", "root")
	os.Setenv("MONGODB_PWD", "root")
	os.Setenv("MONGODB_ADDR", "10.100.175.216:27017")
	os.Setenv("MONGODB_DB", "test")
	os.Setenv("MONGODB_TIMEOUT", "10")
	os.Setenv("MONGODB_POOLSIZE", "30")

	_, err := utils.GetMongoDb("movie")
	if err != nil {
		t.Errorf("%v", err)
	}
}
