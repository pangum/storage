package param

import (
	"time"

	"gitlab.com/ruijc/storage/core"
)

type Url struct {
	Type    core.UrlType
	Export  core.Export
	Expires time.Duration
}

func NewUrl() *Url {
	return &Url{
		Type:    core.UrlType_URL_TYPE_INLINE,
		Export:  core.Export_EXPORT_INTERNET,
		Expires: 3 * time.Hour,
	}
}
