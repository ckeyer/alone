package conf

import (
	"testing"
)

//
func TestLoadConfig(t *testing.T) {
	file := "./v1.json"
	c := GetConfig(file)
	if c == nil {
		t.Error("Get nil config")
	}
}
