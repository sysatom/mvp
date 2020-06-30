package main

import (
	"fmt"
	"github.com/sysatom/mvp/server/controller"
	"github.com/valyala/fasthttp"
)

type ServerHandler struct {
	foobar string
}

func (h *ServerHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		_, _ = fmt.Fprintf(ctx, "/")
	case "/auth":
		controller.AuthHandler(ctx)
	case "/goods":
		controller.GoodsHandler(ctx)
	case "/search":
		controller.SearchHandler(ctx)
	case "/cart":
		controller.CartHandler(ctx)
	case "/order":
		controller.OrderHandler(ctx)
	case "/pay":
		controller.PayHandler(ctx)
	}
}

func main() {
	serverHandler := &ServerHandler{
		foobar: "foobar",
	}
	err := fasthttp.ListenAndServe(":8088", serverHandler.HandleFastHTTP)
	if err != nil {
		fmt.Println(err)
	}
}
