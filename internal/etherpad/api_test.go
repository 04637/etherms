package etherpad

import (
	"aid.dev/etherms/config"
	"testing"
)

func TestInitPad(t *testing.T) {
	conf := config.Load()
	InitPad(conf)
}

func TestListPad(t *testing.T) {
	conf := config.Load()
	InitPad(conf)
	padIds := ListAllPads()
	t.Log(padIds)
	t.Log(padIds[0])
}

func TestListMyPads(t *testing.T) {
	TestInitPad(t)
	ListMyPads("X5615")
}

func TestNewUser(t *testing.T) {
	TestInitPad(t)
	id := NewUser("X5615", "籍观通")
	t.Log(id)
}
