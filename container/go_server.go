package main

import (
	// "fmt"
	// "math"
	// "math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.GET("/", func(c echo.Context) error {
		/* x := 1 + rand.Float64()
		str := ""
		for i := 0; i < 2000000; i++ {
			x = math.Tan(x)
			s := fmt.Sprintf("%f", x)
			str += s
		} */

		return c.HTML(http.StatusOK, "Hello, assignment 2 Bikash and Utkarsh! <3" /* + str */)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
