package controllers

import (
	"NoJS_codefest_server/functions"
	"NoJS_codefest_server/models"
	"NoJS_codefest_server/services"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

func UserRegisterController(c echo.Context) error {
	//[1. Definir los datos x defecto para un usuario && parsear datos de la request]
	var userData models.User
	userData.Us_status = true
	userData.Ty_id = 2 //Rol usuario x defecto
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Cargar la imagen de perfil al servidor]
	if image_data, err := c.FormFile("avatar"); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request: Error uploading image"})
	} else {
		if image_name, err := functions.UploadImageFunction(image_data, []string{".png", ".jpeg", ".webp"}); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error processing image"})
		} else {
			userData.Us_avatar = image_name
		}
	}
	//[3. Crear el registro en la base de datos]
	if err := services.UserRegisterService(userData); err != nil {
		functions.DeleteImageHandler(userData.Us_avatar)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error: Error registering in database"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

func UserLoginController(c echo.Context) error {
	//[1. Parsear datos de la request]
	email := c.QueryParam("email")
	pwd := c.QueryParam("password")
	if email == "" || pwd == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message:": "Bad Request: Invalid parameters"})
	}
	//[2. Consultar el usuario en la base de datos]
	userData, err := services.UserLoginService(email, pwd)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message:": "Not Found: Invalid user credentials"})
	}
	//[3. Crear token de sesión]
	userData.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userData)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}
	userData.Token = t
	return c.JSON(http.StatusOK, userData)
}
