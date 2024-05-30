package gee

import (
	"net/http"
)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: &router{
		root:     map[string]*node{},
		handlers: map[string]HandlerFunc{},
	}}
}

func (g *Engine) addRoute(method, param string, handler HandlerFunc) {
	g.router.addRoute(method, param, handler)
}

func (g *Engine) Get(param string, handler HandlerFunc) {
	g.addRoute("GET", param, handler)
}
func (g *Engine) Post(param string, handler HandlerFunc) {
	g.addRoute("POST", param, handler)
}

func (g *Engine) Run(port string) {
	http.ListenAndServe(":"+port, g)
}

func (g *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	g.router.handle(c)
}
