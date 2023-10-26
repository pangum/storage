package get

import (
	"gitea.com/ruijc/app"
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/pangum/grpc"
	"github.com/pangum/pangu"
)

type Client struct {
	pangu.Get

	Client *grpc.Client
	Http   *http.Client
	Logger log.Logger
	App    app.Id
}
