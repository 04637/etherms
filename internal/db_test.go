package internal

import (
	"aid.dev/etherms/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	conf := config.Load()
	InitDB(conf)
}

func TestMigrateSchema(t *testing.T) {
	TestInitDB(t)
	migrateSchema()
}
