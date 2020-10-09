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
	pad = etherpadlite.NewEtherpadLite(conf.API.APIKey)
	pad.BaseURL = conf.API.BaseURL
	pad.RaiseEtherpadErrors = true
}

// ListAllPads 获取所有文本列表
func ListAllPads() []string {
	resp, err := pad.ListAllPads(context.Background())
	internal.HandleErr(err)
	a := resp.Data["padIDs"].([]interface{})
	padIDs := make([]string, len(a))
	for i, v := range a {
		padIDs[i] = v.(string)
	}
	return padIDs
}

// ListMyhPads 获取我的所有文本
func ListMyPads(userId string) []string {
	resp, err := pad.ListPadsOfAuthor(context.Background(), "12")
	internal.HandleErr(err)
	fmt.Println(resp.Data)
	return nil
}

// NewUser 创建用户
func NewUser(id, name string) string {
	resp, err := pad.CreateAuthorIfNotExistsFor(context.Background(), id, name)
	internal.HandleErr(err)
	authorId := resp.Data["authorID"].(string)
	return authorId
}
