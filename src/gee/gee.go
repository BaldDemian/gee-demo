package gee

import (
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// method: GET, POST...
// pattern: regex expression
// handler: specific HandleFunc
func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET add a GET method handler
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST add a POST handler
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

// ServeHTTP 拦截所有请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// Run 包装http.ListenAndServe
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
