package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitRouters(app *fiber.App) {
	app.Get("/health", controller.Health)

	initSwagger(app)
	initApi(app)
}

func initApi(app *fiber.App) {
	v1 := app.Group("/api/v1")

	initUserRouters(v1)
	initWalletRouters(v1)
	initAuthRouters(v1)
	initTransactionRouters(v1)
}

func initSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:8080/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}
