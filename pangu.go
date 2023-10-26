package storage

import (
	"gitea.com/ruijc/storage/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	creator := new(plugin.Creator)
	pangu.New().Get().Dependency().Put(
		creator.NewClient,
	).Build().Build().Apply()
}
