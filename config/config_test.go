package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf := Load()
	if conf.API.APIKey == "" {
		t.Error("apiKey is nil")
	}
	if conf.API.BaseURL == "" {
		t.Error("baseUrl is nil")
	}
}
