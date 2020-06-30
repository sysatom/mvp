package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func AuthHandler(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "/auth")
}

func GoodsHandler(ctx *fasthttp.RequestCtx) {

}

func SearchHandler(ctx *fasthttp.RequestCtx) {

}

func CartHandler(ctx *fasthttp.RequestCtx) {

}

func OrderHandler(ctx *fasthttp.RequestCtx) {

}

func PayHandler(ctx *fasthttp.RequestCtx) {

}
