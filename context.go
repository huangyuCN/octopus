package octopus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string //request path
	Method     string //request method
	StatusCode int    // response code
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// TextPlain 以 text/plain 作为 content-Type 返回
func (c *Context) TextPlain(code int, format string, values ...interface{}) error {
	c.SetHeader("content-Type", "text/plain")
	c.Status(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	return err
}

func (c *Context) Json(code int, object interface{}) error {
	c.SetHeader("content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(object); err != nil {
		http.Error(c.Writer, err.Error(), 500) //服务器错误
		return err
	}
	return nil
}

func (c *Context) Data(code int, data []byte) error {
	c.Status(code)
	_, err := c.Writer.Write(data)
	return err
}

func (c *Context) Html(code int, html string) error {
	c.Status(code)
	c.SetHeader("content-Type", "text/html")
	_, err := c.Writer.Write([]byte(html))
	return err
}
