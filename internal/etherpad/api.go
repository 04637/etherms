package etherpad

import (
	"aid.dev/etherms/config"
	"aid.dev/etherms/internal"
	"context"
	"fmt"
	"github.com/FabianWe/etherpadlite-golang"
)

var pad *etherpadlite.EtherpadLite

// InitPad init a etherpad api
func InitPad(conf config.Conf) {
	pad = etherpadlite.NewEtherpadLite(conf.Api.ApiKey)
	pad.BaseURL = conf.Api.BaseUrl
	pad.RaiseEtherpadErrors = true
}

func ListPad() []string {
	resp, err := pad.ListAllPads(context.Background())
	internal.HandleErr(err)
	a := resp.Data["padIDs"].([]interface{})
	padIDs := make([]string, len(a))
	for i, v := range a {
		padIDs[i] = v.(string)
	}
	return padIDs
}

func ListMyPads() []string {
	resp, err := pad.PadUsers(context.Background(), "12")
	internal.HandleErr(err)
	fmt.Println(resp.Data)
	return nil
}

func GetAuthorName(authorId string) string {
	resp, err := pad.GetAuthorName(context.Background(), authorId)
	internal.HandleErr(err)
	fmt.Println(resp.Data)
	return ""
}
