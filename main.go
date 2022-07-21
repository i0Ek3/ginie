package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"ginie"
)

type Test struct {
	Name string
	Code string
}

func main() {
	r := ginie.Default()
	r.GET("/", func(c *ginie.Context) {
		c.String(http.StatusOK, "Hi there.\n")
	})

	r.GET("/panic", func(c *ginie.Context) {
		names := []string{"i0Ek3"}
		// simulation panic
		c.String(http.StatusOK, names[100])
	})

	// template rendering and static resource service
	{
		FormatAsDate := func(t time.Time) string {
			y, m, d := t.Date()
			return fmt.Sprintf("%d-%02d-%02d", y, m, d)
		}

		r.SetFuncMap(template.FuncMap{
			"FormatAsDate": FormatAsDate,
		})
		r.LoadHTMLGlob("templates/*")
		// Static maps ./static to /assets
		r.Static("/assets", "./static")
	}

	v0 := r.Group("/v0")
	{
		t1 := &Test{Name: "fishcat", Code: "01"}
		t2 := &Test{Name: "catfish", Code: "10"}
		v0.GET("/", func(c *ginie.Context) {
			c.HTML(http.StatusOK, "css.tmpl", nil)
		})

		v0.GET("/test", func(c *ginie.Context) {
			c.HTML(http.StatusOK, "arr.tmpl", ginie.H{
				"title": "ginie",
				"ts":    [2]*Test{t1, t2},
			})
		})
		v0.GET("/date", func(c *ginie.Context) {
			c.HTML(http.StatusOK, "custom_func.tmpl", ginie.H{
				"title": "ginie",
				"now":   time.Date(2022, 16, 10, 0, 0, 0, 0, time.UTC),
			})
		})
	}

	v1 := r.Group("/v1")
	{
		r.GET("/index", func(c *ginie.Context) {
			c.HTML(http.StatusOK, "", "<h1>Index Page.</h1>")
		})

		v1.GET("/", func(c *ginie.Context) {
			c.HTML(http.StatusOK, "", "<h1>Hi there, this is Ginie!</h1>")
		})

		v1.GET("/hello", func(c *ginie.Context) {
			c.String(http.StatusOK, "Hi %s, this is a test at route %s.\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(applyForV2())
	{
		v2.GET("/hello/:name", func(c *ginie.Context) {
			c.String(http.StatusOK, "Hi %s, this is a test a route %s.\n", c.Param("name"), c.Path)
		})

		v2.GET("/assets/*filepath", func(c *ginie.Context) {
			c.JSON(http.StatusOK, ginie.H{"filepath": c.Param("filepath")})
		})

		v2.POST("/login", func(c *ginie.Context) {
			c.JSON(http.StatusOK, ginie.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":8888")
}

func applyForV2() ginie.HandlerFunc {
	return func(c *ginie.Context) {
		t := time.Now()
		c.Fail(200, "OK")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
