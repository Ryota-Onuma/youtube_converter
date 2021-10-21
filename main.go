package main

import (
  "log"
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"

)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/health_check", health_check)
  e.POST("/mp3_converter", mp3_converter)

  // Static
  e.Static("/", "./templates/index.html")

  // Start server
  e.Logger.Fatal(e.Start(":3000"))
}

// Handler
func health_check(c echo.Context) error {
  return c.String(http.StatusOK, "OK")
}



func mp3_converter(c echo.Context) error{
  url := c.FormValue("url")

  log.Print("----------")
  log.Print(url)
  log.Print("----------")

  return c.String(http.StatusOK, url)
}
