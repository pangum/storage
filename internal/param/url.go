package param

import (
	"time"

	"gitlab.com/ruijc/storage/core/kernel"
)

type Url struct {
	Type    kernel.UrlType
	Export  kernel.Export
	Expires time.Duration
}

func NewUrl() *Url {
	return &Url{
		Type:    kernel.UrlType_URL_TYPE_INLINE,
		Export:  kernel.Export_EXPORT_INTERNET,
		Expires: 3 * time.Hour,
	}
}
