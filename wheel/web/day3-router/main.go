package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.Get("/", func(c *gee.Context) {
		fmt.Println("into index")
		c.HTML(http.StatusOK, "<h1>gee!</h1>")
	})
	engine.Get("/hello", func(c *gee.Context) {
		// expect /hello?name=jxy
		c.String(http.StatusOK, "hello %v", c.Query("name"))
	})
	engine.Get("/hello/:name", func(c *gee.Context) {
		// expect /hello?name=jxy
		c.String(http.StatusOK, "hello %v", c.Param("name"))
	})
	engine.Get("/asserts/*filePath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"filePath": c.Param("filePath"),
		})
	})

	engine.Run("9999")
}
