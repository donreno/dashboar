package main

import (
	"log"
	"time"

	"github.com/donreno/dashboar/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	log.Println("Main!!!")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./assets")

	app.Use(limiter.New(limiter.Config{
		Max:               10,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cache.New(cache.Config{
		Expiration: 30 * time.Second,
	}))

	dashboardService := service.New()

	app.Get("/", func(c *fiber.Ctx) error {

		dashboard := dashboardService.LoadDashboard()

		return c.Render("index", dashboard)
	})

	log.Fatal(app.Listen(":3000"))
}
