package octopus

import (
	"fmt"
	"testing"
)

func TestEngine_RUN(t *testing.T) {
	engine := New()
	engine.GET("/get", func(ctx *Context) {
		query := ctx.Query("param")
		err := ctx.TextPlain(200, "%s", query)
		if err != nil {
			t.Fatalf("response %s failed, error:%s", ctx.Path, err)
		}
	})
	engine.POST("/post", func(ctx *Context) {
		data := ctx.PostForm("param")
		err := ctx.TextPlain(200, "%s", data)
		if err != nil {
			t.Fatalf("response %s failed, error:%s", ctx.Path, err)
		}
	})
	err := engine.RUN("127.0.0.1", "8888")
	if err != nil {
		t.Fatalf("server run error:%s \n", err)
		return
	}
	defer func(engine *Engine) {
		err := engine.STOP()
		if err != nil {
			t.Fatalf("stop http server failed, error:%s", err)
		}
		fmt.Println("http server stopped")
	}(engine)
}
