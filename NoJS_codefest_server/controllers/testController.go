package controllers

import (
	"NoJS_codefest_server/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func TestController(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(*models.User)
	return c.JSON(http.StatusOK, claims)
}
