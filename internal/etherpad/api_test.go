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
	padIds := etherpad.ListAllPads()
	t.Log(padIds)
	t.Log(padIds[0])
}

func TestListMyPads(t *testing.T) {
	TestInitPad(t)
	etherpad.ListMyPads("X5615")
}

func TestNewUser(t *testing.T) {
	TestInitPad(t)
	id := etherpad.NewUser("X5615", "籍观通")
	t.Log(id)
}
