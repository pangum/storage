package plugin

import (
	"github.com/pangum/storage/internal/internal/constant"
	"github.com/pangum/storage/internal/put"
	"gitlab.com/ruijc/storage/protocol"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewClient(client put.Client) *protocol.FxClient {
	return protocol.NewFxClient(client.Client.Connection(constant.LabelStorage), client.Http.Client, client.Logger)
}
