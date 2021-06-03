package octopus

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPanic(t *testing.T) {
	r := Default()
	r.GET("/", func(c *Context) {
		c.TextPlain(http.StatusOK, "Hello octopus\n")
	})
	r.GET("/panic", func(c *Context) {
		names := []string{"octopus"}
		c.TextPlain(http.StatusOK, names[100])
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
