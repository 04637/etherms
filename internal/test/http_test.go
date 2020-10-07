package test

import (
	"testing"

	"aid.dev/etherms/internal"
)

func TestGet(t *testing.T) {
	params := internal.NewParams().
		Add("hello", "world").
		Add("hello1", "wold")
	for k, v := range params.Data() {
		t.Logf("%s, %s", k, v)
	}
}
