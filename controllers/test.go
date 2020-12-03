package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type TestController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @router /test/index [get]
func (c *TestController) Get() {
	c.TplName = "danmuDemo.html"
}

func (c *TestController) WsFunc() {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err = conn.WriteMessage(websocket.TextMessage, []byte("Hello")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
