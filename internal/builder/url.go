package builder

import (
	"time"

	"github.com/pangum/storage/internal/getter"
	"github.com/pangum/storage/internal/param"
	"gitlab.com/ruijc/storage/core/kernel"
	"gitlab.com/ruijc/storage/protocol"
)

type Url struct {
	client *protocol.FxClient
	param  *param.Url
}

func NewUrl(client *protocol.FxClient) *Url {
	return &Url{
		client: client,
		param:  param.NewUrl(),
	}
}

func (u *Url) Inline() (url *Url) {
	u.param.Type = kernel.UrlType_URL_TYPE_INLINE
	url = u

	return
}

func (u *Url) Attachment() (url *Url) {
	u.param.Type = kernel.UrlType_URL_TYPE_ATTACHMENT
	url = u

	return
}

func (u *Url) Type(typ kernel.UrlType) (url *Url) {
	u.param.Type = typ
	url = u

	return
}

func (u *Url) Internet() (url *Url) {
	u.param.Export = kernel.Export_EXPORT_INTERNET
	url = u

	return
}

func (u *Url) Export(export kernel.Export) (url *Url) {
	u.param.Export = export
	url = u

	return
}

func (u *Url) Expires(expires time.Duration) (url *Url) {
	u.param.Expires = expires
	url = u

	return
}

func (u *Url) Build() *getter.Url {
	return getter.NewUrl(u.client, u.param)
}
