package main

import (
	"github.com/Juminiy/kube/cmd/clipboard/clipboard_fast/handler"
	"github.com/gofiber/fiber/v3"
)

func RESTAPI(app *fiber.App) {
	app.Get("/", handler.ClipAddAndList)
	app.Get("/add", handler.ClipAddAndList)
	app.Get("/list", handler.ClipAddAndList)
	app.Get("/search", handler.ClipSearch)
	app.Get("/del", handler.ClipDelAndList)
}
