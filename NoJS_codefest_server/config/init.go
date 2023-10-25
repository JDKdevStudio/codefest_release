package config

import (
	"NoJS_codefest_server/routes"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartService() {
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch},
	}))
	routes.DefineRoutes(server)
	if err := server.Start(":80"); err != nil {
		log.Fatal(err)
	}
}
