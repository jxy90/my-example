package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	//request info
	Path   string
	Method string
	Params map[string]string
	//response info
	StatusCode int
	//middleware
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		//StatusCode: 0,
		index: -1,
	}
}

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	//return c.Req.Form.Get(key)
	return c.Req.FormValue(key)
}

func (c *Context) Status(statusCode int) {
	c.StatusCode = statusCode
	c.Writer.WriteHeader(statusCode)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Data(statusCode int, data []byte) {
	c.Writer.Write(data)
	c.Status(statusCode)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	str := fmt.Sprintf(format, values...)
	c.Status(code)
	c.Writer.Write([]byte(str))
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
