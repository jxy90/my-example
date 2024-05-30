package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: map[string]HandlerFunc{}}
}

func (r *Router) addRoute(method, param string, handler HandlerFunc) {
	key := method + "-" + param
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "404 Not Found %v \n", key)
	}
}
