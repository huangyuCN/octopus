package octopus

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"
	"time"
)

type Student struct {
	Name string
	Age  int8
}

func FormatAsData(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func TestTemple(t *testing.T) {
	r := New()
	r.Use(Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsData,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	r.GET("/", func(c *Context) {
		c.Html(http.StatusOK, "css.tmpl", H{})
	})
	stu1 := &Student{Name: "octopus", Age: 10}
	stu2 := &Student{Name: "hack", Age: 22}
	r.GET("/students", func(c *Context) {
		c.Html(http.StatusOK, "arr.tmpl", H{
			"title":  "octopus",
			"stuArr": [2]*Student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *Context) {
		c.Html(http.StatusOK, "custom_func.tmpl", H{
			"title": "octopus",
			"now":   time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC),
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
