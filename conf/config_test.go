package conf

import (
	"testing"
)

//
func TestLoadConfig(t *testing.T) {
	file := "v1_test.conf"
	c := GetConfig(file)
	if c == nil {
		t.Error("Get nil config")
	}
}
