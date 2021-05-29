package octopus

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEngine_RUN(t *testing.T) {
	engine := New()
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello back"))
		if err != nil {
			t.Fatalf("response %s failed, error:%s", r.URL.Path, err)
		}
	}
	engine.GET("/get", handler)
	engine.POST("/post", handler)
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
