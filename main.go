package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"

  "youtube_converter/download"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/health_check", health_check)
  e.POST("/mp4_converter", mp4_converter)

  // Static
  e.Static("/", "./templates/index.html")

  // Start server
  e.Logger.Fatal(e.Start(":3000"))
}

func health_check(c echo.Context) error {
  return c.String(http.StatusOK, "OK")
}


func mp4_converter(c echo.Context) error{
  youtube_url := c.FormValue("youtube_url")
  download.Mp4_download(youtube_url)
  return c.String(http.StatusOK, youtube_url)
}
