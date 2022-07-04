package webapp

import (
	"fmt"
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
	"github.com/donreno/dashboar/internal/types"
)

func InitWebApp(getDashBOAR service.DashboardRetriever, conf *types.EnvConfig) error {

	log.Println("Initializing Web App 🌎 🐗")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(compress.New())

	app.Use(limiter.New(limiter.Config{
		Max:               conf.LimiterMaxRequests,
		Expiration:        time.Duration(conf.LimiterExpireTimeSeconds) * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(cache.New(cache.Config{
		Expiration: time.Duration(conf.CacheExpireTimeSeconds) * time.Second,
	}))

	app.Static("/assets", "./assets")

	app.Use(favicon.New(
		favicon.Config{
			File: "./assets/favicon.ico",
		},
	))

	app.Get("/", func(c *fiber.Ctx) error {

		dashBOAR, err := getDashBOAR()

		if err != nil {
			log.Fatalf("Error getting dashBOAR %q", err.Error())
			c.Render("error", err)
		}

		return c.Render("index", dashBOAR)
	})

	log.Println("Web App initialized! 🌎🍻🐗")

	return app.Listen(fmt.Sprintf(":%d", conf.Port))
}
