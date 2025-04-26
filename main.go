package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
  c.Links(
    "https://github.com/kmscom/Browser-Cache-Folder-Changer", "Web Browser Cache Folder Changer",
    "https://github.com/kmscom/Browser-Cache-Folder-Changer/releases", "Download",
  )
})

	app.Listen(getPort())
}
