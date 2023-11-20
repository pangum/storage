package core

import (
	"context"
	"time"

	"github.com/goexl/gox"
	"github.com/pangum/storage/internal/builder"
	"gitlab.com/ruijc/storage/core/file"
	"gitlab.com/ruijc/storage/core/kernel"
	"gitlab.com/ruijc/storage/protocol"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Client struct {
	client *protocol.FxClient

	_ gox.CannotCopy
}

func NewClient(client *protocol.FxClient) *Client {
	return &Client{
		client: client,
	}
}

func (s *Client) Initiate(
	ctx context.Context,
	mime string, name string,
	parts int32, start int32,
	expires time.Duration,
) (*file.InitiateRsp, error) {
	return s.client.Initiate(ctx, &file.InitiateReq{
		Mime:    mime,
		Name:    name,
		Parts:   parts,
		Start:   start,
		Expires: durationpb.New(expires),
	})
}

func (s *Client) Put(ctx context.Context, name string, content []byte) (int64, error) {
	return s.client.Put(ctx, name, content)
}

func (s *Client) Url() *builder.Url {
	return builder.NewUrl(s.client)
}

func (s *Client) Complete(ctx context.Context, id int64, parts *[]*kernel.Part) (*file.CompleteRsp, error) {
	return s.client.Complete(ctx, &file.CompleteReq{
		Id:    id,
		Parts: *parts,
	})
}

func (s *Client) Abort(ctx context.Context, id int64) (*file.AbortRsp, error) {
	return s.client.Abort(ctx, &file.AbortReq{
		Id: id,
	})
}

func (s *Client) Use(ctx context.Context, id int64) (*file.UseRsp, error) {
	return s.client.Use(ctx, &file.UseReq{
		Id: id,
	})
}

func (s *Client) Release(ctx context.Context, id int64) (*file.ReleaseRsp, error) {
	return s.client.Release(ctx, &file.ReleaseReq{
		Id: id,
	})
}

func (s *Client) Delete(ctx context.Context, id int64) (*file.DeleteRsp, error) {
	return s.client.Delete(ctx, &file.DeleteReq{
		Id: id,
	})
}
