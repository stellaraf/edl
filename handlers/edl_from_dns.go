package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/stellaraf/edl/lib/doh"
)

const cache_header_value string = "s-maxage=3600, stale-while-revalidate=14400"

func createList(response *doh.DOHResponse) string {
	list := []string{}
	for _, answer := range response.Answer {
		if answer.Type == 1 || answer.Type == 28 {
			list = append(list, answer.Data)
		}
	}
	return strings.Join(list, "\n")
}

func edlFromDNS(ctx *fiber.Ctx, name string) error {
	client := doh.New()
	query, err := client.A(name)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	list := createList(query)
	ctx.Set("content-type", "text/plain")
	ctx.Set("cache-control", cache_header_value)
	return ctx.Status(200).Send([]byte(list))
}

func EDLFromDNSHandler(query string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return edlFromDNS(ctx, query)
	}
}
