package tg

import (
	_ "embed"
	"github.com/valyala/fasthttp"
)

const (
	ApiScheme = "https"
	ApiHost   = "api.telegram.org"
)

var (
	//go:embed token.txt
	Token string

	DefaultUri = ApiScheme + "://" + ApiHost + "/bot" + Token + "/"

	ApiClient = fasthttp.HostClient{Addr: ApiHost, IsTLS: true}
)
