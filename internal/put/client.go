package put

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/pangum/grpc"
	"github.com/pangum/pangu"
)

type Client struct {
	pangu.Put

	Client *grpc.Client
	Http   *http.Client
	Logger log.Logger
}
