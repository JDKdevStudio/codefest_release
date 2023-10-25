package controllers

import (
	"NoJS_codefest_server/functions"
	"NoJS_codefest_server/models"
	"NoJS_codefest_server/services"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ProjectRegisterController(c echo.Context) error {
	//[1. Definir los datos x defecto para un proyecto && parsear request]
	claims := c.Get("user").(*jwt.Token).Claims.(*models.User)
	var projectData models.Project
	projectData.Pr_status = true
	if err := c.Bind(&projectData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Cargar la imagen de perfil al servidor]
	if image_banner, err := c.FormFile("banner"); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request: Error uploading image"})
	} else {
		if image_name, err := functions.UploadImageFunction(image_banner, []string{".png", ".jpeg", ".webp"}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error processing image"})
		} else {
			projectData.Pr_banner = image_name
		}
	}
	//[3. Crear el registro en la base de datos]
	if err := services.ProjectRegisterService(projectData, claims.Us_id); err != nil {
		functions.DeleteImageHandler(projectData.Pr_banner)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error registering in database"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

func ProyectListController(c echo.Context) error {
	var res []models.Project
	res, err := services.ProjectGetListService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error retrieving from database"})
	}
	return c.JSON(http.StatusOK, res)
}
