package builder

import (
	"time"

	"gitea.com/ruijc/storage/internal"
	"gitea.com/ruijc/storage/internal/getter"
	"gitea.com/ruijc/storage/internal/param"
	"gitlab.com/ruijc/storage/core"
)

type Url struct {
	client *internal.Client
	param  *param.Url
}

func NewUrl(client *internal.Client) *Url {
	return &Url{
		client: client,
		param:  param.NewUrl(),
	}
}

func (u *Url) Inline() (url *Url) {
	u.param.Type = core.UrlType_URL_TYPE_INLINE
	url = u

	return
}

func (u *Url) Attachment() (url *Url) {
	u.param.Type = core.UrlType_URL_TYPE_ATTACHMENT
	url = u

	return
}

func (u *Url) Internet() (url *Url) {
	u.param.Export = core.Export_EXPORT_INTERNET
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
