package octopus

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestEngine_RUN(t *testing.T) {
	r := New()
	r.Static("/assets", "./static")
	v1 := r.Group("/v1")
	v1.GET("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.TextPlain(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/hello/:name", func(c *Context) {
		// expect /hello/geektutu
		c.TextPlain(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v2.POST("/login", func(c *Context) {
		c.Json(http.StatusOK, H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
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

func onlyForV2() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Fail(500, "Internal server error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
