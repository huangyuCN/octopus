package octopus

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEngine_RUN(t *testing.T) {
	r := New()
	r.GET("/", func(c *Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.TextPlain(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *Context) {
		// expect /hello/geektutu
		c.TextPlain(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *Context) {
		c.Json(http.StatusOK, H{"filepath": c.Param("filepath")})
	})
	err := r.RUN("127.0.0.1", "8888")
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
	}(r)
}
