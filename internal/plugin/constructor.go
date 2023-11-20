package plugin

import (
	"github.com/pangum/storage/internal/get"
	"github.com/pangum/storage/internal/internal/constant"
	"gitlab.com/ruijc/storage/protocol"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewClient(get get.Client) *protocol.FxClient {
	return protocol.NewFxClient(get.Client.Connection(constant.LabelStorage), get.Http.Client, get.Logger)
}
