package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/index", func(c *gee.Context) {
		fmt.Println("into index")
		c.HTML(http.StatusOK, "<h1>index page!</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>v1 page!</h1>")
		})
		v1.Get("/hello", func(c *gee.Context) {
			// expect /hello?name=jxy
			c.String(http.StatusOK, "hello v1 %v", c.Query("name"))
		})
	}
	v2 := r.Group("/v2")
	{
		v2.Get("/hello/:name", func(c *gee.Context) {
			// expect /hello?name=jxy
			c.String(http.StatusOK, "hello %v", c.Param("name"))
		})
		v2.Get("/asserts/*filePath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"filePath": c.Param("filePath"),
			})
		})
	}

	r.Run("9999")
}
