package main

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	e.POST("/upload", Upload())
	host, _ := os.Hostname()
	log.Println("access http://", host)

	e.Run(standard.New(":1323"))
	//log.Println("access to http://localhst:1323")
}
