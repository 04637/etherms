package internal

import (
	"testing"
)

func TestGet(t *testing.T) {
	params := NewParams().
		Add("ie", "UTF-8").
		Add("wd", "1212")
	Get("https://wwww.github.com", params)
}
