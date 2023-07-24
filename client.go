package storage

import (
	"context"
	"time"

	"gitea.com/ruijc/app"
	"gitea.com/ruijc/storage/internal"
	"gitea.com/ruijc/storage/internal/builder"
	"github.com/goexl/gox"
	"gitlab.com/ruijc/storage/core"
	"gitlab.com/ruijc/storage/file"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Storage 文件存储客户端
type Storage struct {
	client *internal.Client
	app    app.Id

	_ gox.CannotCopy
}

func newStorage(client *internal.Client, app app.Id) *Storage {
	return &Storage{
		client: client,
		app:    app,
	}
}

func (s *Storage) Initiate(
	ctx context.Context,
	mime string, name string,
	parts int32, start int32,
	expires time.Duration,
) (*file.InitiateRsp, error) {
	return s.client.Initiate(ctx, &file.InitiateReq{
		App:     int64(s.app),
		Mime:    mime,
		Name:    name,
		Parts:   parts,
		Start:   start,
		Expires: durationpb.New(expires),
	})
}

func (s *Storage) Put(ctx context.Context, name string, content []byte) (int64, error) {
	return s.client.Put(ctx, int64(s.app), name, content)
}

func (s *Storage) Url() *builder.Url {
	return builder.NewUrl(s.client)
}

func (s *Storage) Complete(ctx context.Context, id int64, parts []*core.Part) (*file.CompleteRsp, error) {
	return s.client.Complete(ctx, &file.CompleteReq{
		Id:    id,
		Parts: parts,
	})
}

func (s *Storage) Abort(ctx context.Context, id int64) (*file.AbortRsp, error) {
	return s.client.Abort(ctx, &file.AbortReq{
		Id: id,
	})
}

func (s *Storage) Delete(ctx context.Context, id int64) (*file.DeleteRsp, error) {
	return s.client.Delete(ctx, &file.DeleteReq{
		Id: id,
	})
}

func (s *Storage) Deletes(ctx context.Context, ids []int64) (*file.DeletesRsp, error) {
	return s.client.Deletes(ctx, &file.DeletesReq{
		Ids: ids,
	})
}
