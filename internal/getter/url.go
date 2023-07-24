package getter

import (
	"context"

	"gitea.com/ruijc/storage/internal"
	"gitea.com/ruijc/storage/internal/param"
	"gitlab.com/ruijc/storage/file"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Url struct {
	client *internal.Client
	param  *param.Url
}

func NewUrl(client *internal.Client, param *param.Url) *Url {
	return &Url{
		client: client,
		param:  param,
	}
}

func (u *Url) Get(ctx context.Context, id int64) (url string, err error) {
	req := new(file.UrlReq)
	req.Id = id
	req.Type = u.param.Type
	req.Export = u.param.Export
	req.Expires = durationpb.New(u.param.Expires)
	if rsp, ue := u.client.Url(ctx, req); nil != ue {
		err = ue
	} else {
		url = rsp.Url
	}

	return
}

func (u *Url) Gets(ctx context.Context, ids []int64) (url map[int64]string, err error) {
	req := new(file.UrlsReq)
	req.Ids = ids
	req.Type = u.param.Type
	req.Export = u.param.Export
	req.Expires = durationpb.New(u.param.Expires)
	if rsp, ue := u.client.Urls(ctx, req); nil != ue {
		err = ue
	} else {
		url = rsp.Urls
	}

	return
}
