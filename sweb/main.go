package main

import (
	"github.com/sion96994/go/sweb/see"
	"net/http"
)

func main() {
	r := see.New()

	r.GET("/", func(c *see.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Sion</h1>")
	})
	r.GET("/hello", func(c *see.Context) {
		// expect /hello?name=""
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *see.Context) {
		c.JSON(http.StatusOK, see.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":6666")
}