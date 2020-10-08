package etherpad_test

import (
	"aid.dev/etherms/config"
	"aid.dev/etherms/internal/etherpad"
	"testing"
)

func TestInitPad(t *testing.T) {
	conf := config.LoadConfig()
	etherpad.InitPad(conf)
}

func TestListPad(t *testing.T) {
	conf := config.LoadConfig()
	etherpad.InitPad(conf)
	padIds := etherpad.ListPad()
	t.Log(padIds)
	t.Log(padIds[0])
}

func TestListMyPads(t *testing.T) {
	TestInitPad(t)
	etherpad.ListMyPads()
}
