package storage

import (
	"github.com/pangum/pangu"
	"github.com/pangum/storage/internal/plugin"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.NewClient,
	).Build().Build().Apply()
}
