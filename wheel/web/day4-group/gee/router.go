package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type router struct {
	root     map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		root:     map[string]*node{},
		handlers: map[string]HandlerFunc{}}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}
	r.root[method].insert(pattern, parts, 0)

	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	root, ok := r.root[method]
	if !ok {
		return nil, nil
	}
	node := root.search(searchParts, 0)
	if node == nil {
		return nil, nil
	}
	parts := parsePattern(node.pattern)
	kv := make(map[string]string)
	for index, part := range parts {
		if part[0] == ':' {
			kv[part[1:]] = searchParts[index]
		} else if part[0] == '*' && len(part) > 1 {
			kv[part[1:]] = strings.Join(searchParts[index:], "/")
		}
	}
	return node, kv
}

func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		r.handlers[key](c)
	} else {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "404 Not Found %v \n", c.Path)
	}
	//key := c.Method + "-" + c.Path
	//if handler, ok := r.handlers[key]; ok {
	//	handler(c)
	//} else {
	//	c.Status(http.StatusNotFound)
	//	fmt.Fprintf(c.Writer, "404 Not Found %v \n", key)
	//}
}
