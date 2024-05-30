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

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	fmt.Println("Day7")
	r := gee.Default()

	r.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	r.Get("/panic", func(c *gee.Context) {
		panic(c)
	})

	r.Run("9999")
}

func panic(c *gee.Context) {
	names := []string{"12"}
	c.String(http.StatusOK, names[1])
}
