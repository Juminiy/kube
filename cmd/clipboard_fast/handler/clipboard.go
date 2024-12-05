package handler

import (
	"github.com/Juminiy/kube/cmd/clipboard_fast/service"
	"github.com/gofiber/fiber/v3"
)

func ClipAddAndList(c fiber.Ctx) error {
	service.ClipAdd(getQueryv(c))
	return c.JSON(service.ClipList())
}

func ClipDelAndList(c fiber.Ctx) error {
	service.ClipDel(getQueryv(c))
	return c.JSON(service.ClipList())
}

func ClipSearch(c fiber.Ctx) error {
	return c.JSON(service.ClipSearch(getQueryv(c)))
}

func getQueryv(c fiber.Ctx) string {
	v, val, value := c.Query("v"), c.Query("val"), c.Query("value")
	switch {
	case len(v) > 0:
		return v
	case len(val) > 0:
		return val
	case len(value) > 0:
		return val
	default:
		return ""
	}
}
