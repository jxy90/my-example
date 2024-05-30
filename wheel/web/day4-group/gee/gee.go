package gee

import (
	"net/http"
)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix     string
	middleWare []HandlerFunc
	parent     *RouterGroup
	engine     *Engine
}

func New() *Engine {
	engine := &Engine{router: &router{
		root:     map[string]*node{},
		handlers: map[string]HandlerFunc{},
	}}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
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

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		//middleWare: nil,
		parent: group,
		engine: engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method, param string, handler HandlerFunc) {
	//需要加上group前缀
	pattern := group.prefix + param
	group.engine.router.addRoute(method, pattern, handler)
}
func (group *RouterGroup) Get(param string, handler HandlerFunc) {
	group.addRoute("GET", param, handler)
}
func (group *RouterGroup) Post(param string, handler HandlerFunc) {
	group.addRoute("POST", param, handler)
}

func (g *Engine) Run(port string) {
	http.ListenAndServe(":"+port, g)
}

func (g *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	g.router.handle(c)
}
