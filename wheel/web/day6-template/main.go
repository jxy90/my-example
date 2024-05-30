package main

import (
	"fmt"
	"gee"
	"html/template"
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
	fmt.Println("Day6")
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "jxy", Age: 30}
	stu2 := &student{Name: "Jack", Age: 22}
	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Get("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Get("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run("9999")
}
