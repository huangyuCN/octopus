package octopus

import (
	"fmt"
	"net"
	"net/http"
)

// HandlerFunc defines the request handler
// HandlerFunc 定义了http的请求处理方法
type HandlerFunc func(c *Context)

// Engine implement the interface if ServerHttp
type Engine struct {
	ln     *net.Listener
	router *router
}

// New is the constructor of Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加一条路由信息
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method the route of get request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST defines the method the route of post request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//RUN start a http server
func (engine *Engine) RUN(ip string, port string) error {
	address := ip + ":" + port
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("start http server error:%s \n", err)
		return err
	}
	engine.ln = &ln
	return http.Serve(ln, engine)
}

//STOP close the tcp connect
func (engine *Engine) STOP() error {
	ln := *engine.ln
	return ln.Close()
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
