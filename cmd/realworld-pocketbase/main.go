package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/ad-ops/realworld-pocketbase/web/components"
	"github.com/ad-ops/realworld-pocketbase/web/views"
	"github.com/ad-ops/realworld-pocketbase/internal/models"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	use_static := len(os.Getenv("USE_STATIC")) != 0

	if !use_static {
		app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
			e.Router.GET("/", func(c echo.Context) error {
				home := views.Home()
				component := views.App(home)
				component.Render(c.Request().Context(), c.Response().Writer)

				return nil
			}, apis.ActivityLogger(app))
			e.Router.GET("/hello/:name", func(c echo.Context) error {
				name := c.PathParam("name")
				component := views.Hello(name)
				component.Render(c.Request().Context(), c.Response().Writer)

				return nil
			}, apis.ActivityLogger(app))
			e.Router.GET("/hx/tags", func(c echo.Context) error {
				records, err := app.Dao().FindRecordsByExpr("tags")
				if err != nil {
					return err
				}
				tags := make([]string, 0, len(records))
				for _, r := range records {
					tags = append(tags, r.GetString("tag"))
				}
				component := components.Tags(tags)
				component.Render(c.Request().Context(), c.Response().Writer)

				return nil
			}, apis.ActivityLogger(app))
			e.Router.GET("/hx/articles", func(c echo.Context) error {
				// records, err := app.Dao().FindRecordsByExpr("articles")
				// if err != nil {
				// 	return err
				// }
				// tags := make([]string, 0, len(records))
				// for _, r := range records {
				// 	tags = append(tags, r.GetString("tag"))
				// }
				article := models.Article{ Slug: "test" }
				articles := []models.Article{ article }
				component := components.Articles(articles)
				component.Render(c.Request().Context(), c.Response().Writer)

				return nil
			}, apis.ActivityLogger(app))

			return nil
		})
	}

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
