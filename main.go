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
		return c.SendString(`
        <html>
        <body>
            <h1>Web Browser Cache Folder Changer</h1>
            <ul>
                <li><a href="https://github.com/kmscom/Browser-Cache-Folder-Changer">GitHub Repository</a></li>
                <li><a href="https://github.com/kmscom/Browser-Cache-Folder-Changer/releases">Download</a></li>
            </ul>
        </body>
        </html>
    `)
	})

	app.Listen(getPort())
}
