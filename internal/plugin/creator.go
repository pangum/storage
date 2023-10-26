package plugin

import (
	"gitea.com/ruijc/storage/internal/core"
	"gitea.com/ruijc/storage/internal/get"
	"gitlab.com/ruijc/storage/file"
)

type Creator struct {
	// 解决命名空间问题
}

func (c *Creator) NewClient(get get.Client) *core.Client {
	return core.NewClient(file.NewRpxClient(get.Client.Connection("storage"), get.Http.Client, get.Logger), get.App)
}
