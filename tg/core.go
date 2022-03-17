package tg

import "github.com/valyala/fasthttp"

const (
	ApiScheme = "https"
	ApiHost   = "api.telegram.org"
	Token     = "5144931398:AAF65sVvXUMATEj5uhl5BhTYdcTDMJs8KtM"
)

var (
	DefaultUri = ApiScheme + "://" + ApiHost + "/bot" + Token + "/"

	ApiClient = fasthttp.HostClient{Addr: ApiHost, IsTLS: true}
)
