package get

import (
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
}
