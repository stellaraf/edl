package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stellaraf/edl/static"
)

func edlFromStatic(ctx *fiber.Ctx, name string) error {
	list, err := static.GetList(name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	ctx.Set("content-type", "text/plain")
	ctx.Set("cache-control", cache_header_value)
	return ctx.Status(200).Send(list)
}

func EDLFromStaticHandler(query string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return edlFromStatic(ctx, query)
	}
}
