package plugin

import (
	"github.com/pangum/storage/internal/get"
	"github.com/pangum/storage/internal/internal/constant"
	"gitlab.com/ruijc/storage/protocol"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewClient(client get.Client) *protocol.FxClient {
	return protocol.NewFxClient(client.Client.Connection(constant.LabelStorage), client.Http.Client, client.Logger)
}
