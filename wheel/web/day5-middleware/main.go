package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		log.Printf("part1 statusCode:%v URI:%v time:%v v2 only ", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
		//c.String(500, "500 error")
		log.Printf("part2 statusCode:%v URI:%v time:%v v2 only ", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	fmt.Println("Day5")
	r := gee.New()
	r.Use(gee.Logger())
	r.Get("/index", func(c *gee.Context) {
		fmt.Println("into index")
		c.HTML(http.StatusOK, "<h1>index page!</h1>")
	})
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.Get("/hello/:name", func(c *gee.Context) {
			// expect /hello?name=jxy
			c.String(http.StatusOK, "hello %v", c.Param("name"))
		})
	}

	r.Run("9999")
}
