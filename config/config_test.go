package config_test

import (
	"aid.dev/etherms/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf := config.LoadConfig()
	if conf.Api.ApiKey == "" {
		t.Error("apiKey is nil")
	}
	if conf.Api.BaseUrl == "" {
		t.Error("baseUrl is nil")
	}
}
