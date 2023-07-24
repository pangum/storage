package internal

import (
	"gitea.com/ruijc/app"
	"github.com/pangum/grpc"
	"github.com/pangum/http"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
	"gitlab.com/ruijc/storage/file"
)

type (
	Client struct {
		*file.RpxClient
	}

	In struct {
		pangu.In

		Client *grpc.Client
		Http   *http.Client
		Logger logging.Logger
		App    app.Id
	}
)

func NewClient(in In) *Client {
	return &Client{
		RpxClient: file.NewRpxClient(in.Client.Connection("storage"), in.Http.Client, in.Logger),
	}
}
