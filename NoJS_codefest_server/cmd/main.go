package main

import (
	"NoJS_codefest_server/config"
	"NoJS_codefest_server/database"
	"NoJS_codefest_server/env"
)

func main() {
	//Primero cargar las variables de entorno
	env.Init()
	//Precargado del cliente de la base de datos relacional
	database.MySQLClient()
	//Precargado del cliente de la base de datos no relacional
	//Inicialización del servicio
	config.StartService()
}
