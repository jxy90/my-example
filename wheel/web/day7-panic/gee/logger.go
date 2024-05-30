package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		log.Printf("part1 statusCode:%v URI:%v time:%v", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
		log.Printf("part2 statusCode:%v URI:%v time:%v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
