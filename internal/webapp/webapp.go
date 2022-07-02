package webapp

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	"github.com/donreno/dashboar/internal/service"
)

func InitWebApp(getDashBOAR service.DashboardRetriever) error {

	log.Println("Initializing Web App 🌎 🐗")

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

	app.Use(favicon.New())

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cache.New(cache.Config{
		Expiration: 30 * time.Second,
	}))

	app.Get("/", func(c *fiber.Ctx) error {

		dashBOAR, err := getDashBOAR()

		if err != nil {
			log.Fatalf("Error getting dashBOAR %q", err.Error())
			c.Render("error", err)
		}

		return c.Render("index", dashBOAR)
	})

	log.Println("Web App initialized! 🌎🍻🐗")

	return app.Listen(":3000")
}
