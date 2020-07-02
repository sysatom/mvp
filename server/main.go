package main

import (
	"fmt"
	"github.com/sysatom/mvp/server/controller"
	"github.com/sysatom/mvp/server/database"
	"github.com/sysatom/mvp/server/redis"
	"github.com/valyala/fasthttp"
	"os"
)

type ServerHandler struct {
	foobar string
}

func (h *ServerHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		_, _ = fmt.Fprintf(ctx, os.Getenv("APP_NAME"))
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
	// Database
	database.Connect()

	// Redis
	redis.Connect()

	// Http server
	serverHandler := &ServerHandler{
		foobar: "foobar",
	}
	err := fasthttp.ListenAndServe(":5018", serverHandler.HandleFastHTTP)
	if err != nil {
		fmt.Println(err)
	}
}
