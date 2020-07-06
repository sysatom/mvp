package controller

import (
	"encoding/json"
	"fmt"
	"github.com/sysatom/mvp/model"
	"github.com/sysatom/mvp/server/database"
	"github.com/sysatom/mvp/server/redis"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
)

func AuthHandler(ctx *fasthttp.RequestCtx) {
	args := ctx.QueryArgs()
	username := args.Peek("username")
	// ignore
	// password := args.Peek("password")

	k := strings.Builder{}
	k.WriteString("username:")
	for _, i := range username {
		k.WriteByte(i)
	}

	var id int
	val, err := redis.RDB.Get(ctx, k.String()).Result()
	if err != nil {
		fmt.Println(err)

		row := database.DB.QueryRow("select id from mall_user where username = ?", username)

		err := row.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}

		redis.RDB.Set(ctx, k.String(), id, time.Second*60)
	} else {
		id, err = strconv.Atoi(val)
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	data := map[string]interface{}{
		"error": 0,
		"data":  id,
	}
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = ctx.Write(j)
}

func GoodsHandler(ctx *fasthttp.RequestCtx) {
	args := ctx.QueryArgs()
	id, err := args.GetUint("id")
	if err != nil {
		fmt.Println(err)
	}

	k := strings.Builder{}
	k.WriteString("goods:")
	k.WriteString(strconv.Itoa(id))
	var goods model.MallGoods
	val, err := redis.RDB.Get(ctx, k.String()).Result()
	if err != nil {
		fmt.Println(err)

		err := database.DB.QueryRowx("select * from mall_goods where id = ?", id).StructScan(&goods)
		if err != nil {
			fmt.Println(err)
		}
		j, err := json.Marshal(goods)
		if err != nil {
			fmt.Println(err)
		}
		redis.RDB.Set(ctx, k.String(), j, time.Second*60)
	} else {
		err = json.Unmarshal([]byte(val), &goods)
		if err != nil {
			fmt.Println(err)
		}
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	data := map[string]interface{}{
		"error": 0,
		"data":  goods,
	}
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = ctx.Write(j)
}

func SearchHandler(ctx *fasthttp.RequestCtx) {

}

func CartHandler(ctx *fasthttp.RequestCtx) {

}

func OrderHandler(ctx *fasthttp.RequestCtx) {

}

func PayHandler(ctx *fasthttp.RequestCtx) {

}
