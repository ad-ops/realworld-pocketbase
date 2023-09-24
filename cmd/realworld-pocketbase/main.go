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

	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	registry := template.NewRegistry()

	// 	e.Router.GET("/hx/tags", func(c echo.Context) error {
	// 		records, err := app.Dao().FindRecordsByExpr("tags")
	// 		if err != nil {
	// 			return err
	// 		}
	// 		tags := make([]string, 0, len(records))
	// 		for _, r := range records {
	// 			tags = append(tags, r.GetString("tag"))
	// 		}

	// 		html, err := registry.LoadFiles(
	// 			"components/tags.html",
	// 		).Render(tags)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		return c.HTML(http.StatusOK, html)
	// 	}, apis.ActivityLogger(app))

	// 	return nil
	// })

	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	registry := template.NewRegistry()

	// 	e.Router.GET("/", func(c echo.Context) error {
	// 		html, err := registry.LoadFiles(
	// 			"views/layout.html",
	// 			"views/home.html",
	// 		).Render(nil)

	// 		if err != nil {
	// 			return apis.NewNotFoundError("", err)
	// 		}

	// 		return c.HTML(http.StatusOK, html)
	// 	})
	// 	e.Router.GET("/login", func(c echo.Context) error {
	// 		html, err := registry.LoadFiles(
	// 			"views/layout.html",
	// 			"views/login.html",
	// 		).Render(nil)

	// 		if err != nil {
	// 			return apis.NewNotFoundError("", err)
	// 		}

	// 		return c.HTML(http.StatusOK, html)
	// 	})
	// 	e.Router.GET("/register", func(c echo.Context) error {
	// 		html, err := registry.LoadFiles(
	// 			"views/layout.html",
	// 			"views/register.html",
	// 		).Render(nil)

	// 		if err != nil {
	// 			return apis.NewNotFoundError("", err)
	// 		}

	// 		return c.HTML(http.StatusOK, html)
	// 	})
	// 	e.Router.GET("/settings", func(c echo.Context) error {
	// 		html, err := registry.LoadFiles(
	// 			"views/layout.html",
	// 			"views/settings.html",
	// 		).Render(nil)

	// 		if err != nil {
	// 			return apis.NewNotFoundError("", err)
	// 		}

	// 		return c.HTML(http.StatusOK, html)
	// 	})
	// 	e.Router.GET("/article/:slug", func(c echo.Context) error {
	// 		html, err := registry.LoadFiles(
	// 			"views/layout.html",
	// 			"views/register.html",
	// 		).Render(nil)

	// 		if err != nil {
	// 			return apis.NewNotFoundError("", err)
	// 		}

	// 		return c.HTML(http.StatusOK, html)
	// 	})

	// 	return nil
	// })

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
