package main

import (
	"log"
	"github.com/sion96994/go/sweb/see"
	"net/http"
)

func main() {
	r := see.New()
	log.Printf("new r -> %#v", r)

	r.GET("/index", func(c *see.Context) {
		c.HTML(http.StatusOK, "<h1>Index page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *see.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *see.Context) {
			// expect /hello?name=""
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *see.Context) {
			// expect /hello/""
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.GET("/assets/*filepath", func(c *see.Context) {
			c.JSON(http.StatusOK, see.H{"filepath": c.Param("filepath")})
		})
	}

	r.POST("/login", func(c *see.Context) {
		c.JSON(http.StatusOK, see.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":6666")
}