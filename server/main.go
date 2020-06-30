package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type ServerHandler struct {
	foobar string
}

func (h *ServerHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

func main() {
	serverHandler := &ServerHandler{
		foobar: "foobar",
	}
	fasthttp.ListenAndServe(":8088", serverHandler.HandleFastHTTP)
}
