package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/YuheiNakasaka/sayhuuzoku/generator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// Response is response data
type Response struct {
	Name  string `json:"name" xml:"name"`
	Count int    `json:"count" xml:"count"`
}

func handler(c echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(4) + 3
	name, _ := generator.Start(count)
	resp := &Response{
		Name:  name,
		Count: count,
	}
	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()
	e.GET("/", handler)
	e.Run(standard.New(":8080"))
}
