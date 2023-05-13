package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	//在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成
	router *router
}

// New is the constructor of gee.Engine，新建一个Engine对象
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 路由添加
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
// GET请求用于从服务器获取资源
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
// POST请求用于向服务器提交数据
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
