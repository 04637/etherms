package test

import (
	"testing"

	"aid.dev/etherms/internal"
)

func TestGet(t *testing.T) {
	params := internal.NewParams().
		Add("ie", "UTF-8").
		Add("wd", "1212")
	internal.Get("https://wwww.github.com", params)
}
