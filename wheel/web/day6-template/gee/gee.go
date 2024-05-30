package gee

import (
	"html/template"
	"net/http"
	"strings"
)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
	//html template
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

type RouterGroup struct {
	prefix      string
	middleWares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
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

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func (engine *Engine) addRoute(method, param string, handler HandlerFunc) {
	engine.router.addRoute(method, param, handler)
}

func (engine *Engine) Get(param string, handler HandlerFunc) {
	engine.addRoute("GET", param, handler)
}
func (engine *Engine) Post(param string, handler HandlerFunc) {
	engine.addRoute("POST", param, handler)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		//middleWares: nil,
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

func (group *RouterGroup) Use(middleware HandlerFunc) {
	group.middleWares = append(group.middleWares, middleware)
}

func (engine *Engine) Run(port string) {
	http.ListenAndServe(":"+port, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleWares := make([]HandlerFunc, 0)
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middleWares = append(middleWares, group.middleWares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middleWares
	c.engine = engine
	engine.router.handle(c)
}
