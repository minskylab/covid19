package covid19

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/now"
)

type API struct {
	app *fiber.App
}

func NewAPI() (*API, error) {
	app := fiber.New()
	return &API{app: app}, nil
}

func (api *API) routes() {
	g := api.app.Group("/api")

	g.Get("/", func(c *fiber.Ctx) {
		_ = c.JSON(map[string]interface{}{
			"endpoints": map[string]interface{}{
				"/statistics": "Statistics about covid19 virus in Peru",
				"/news":       "News related about covid19 in Peru",
			},
		})
	})

	g.Get("/statistics", func(c *fiber.Ctx) {
		from := c.Query("from")
		to := c.Query("to")

		if from == "" {
			from = now.BeginningOfDay().String()
		}

		if to == "" {
			to = now.EndOfDay().String()
		}

		fromTime, err := now.Parse(from)
		if err != nil {
			_ = c.JSON(map[string]error{"error": err})
			return
		}

		toTime, err := now.Parse(from)
		if err != nil {
			_ = c.JSON(map[string]error{"error": err})
			return
		}

		stats := &CoronaStats{
			From:      fromTime,
			To:        toTime,
			Confirmed: 48,
			Recovered: 0,
			Deaths:    0,
		}

		if err := c.JSON(stats); err != nil {
			_ = c.JSON(map[string]error{"error": err})
			return
		}
	})

	g.Get("/news", func(c *fiber.Ctx) {

	})

}


func (api *API) Run(port int64) error {
	api.routes()
	return api.app.Listen(int(port))
}