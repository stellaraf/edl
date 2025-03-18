package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/stellaraf/edl/handlers"
)

func handler() http.HandlerFunc {
	config := fiber.Config{
		AppName: "EDL",
		Network: "tcp",
	}

	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(requestid.New())

	// EDL Routes
	app.Get("/weave", handlers.EDLFromDNSHandler("allow.us1.weavephone.net"))
	app.Get("/threatlocker/url", handlers.EDLFromStaticHandler("threatlocker_url"))

	return adaptor.FiberApp(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	handler().ServeHTTP(w, r)
}
