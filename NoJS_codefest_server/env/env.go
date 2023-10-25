package env

import (
	"log"
	"os"
)

// Esta función instancia las variables de entorno.
func Init() {
	//Global variables
	set("JWT_SECRET", "9a748514-8668-49f9-a2d4-ee4b3ed00cc7")
	//MySQL variables
	set("MYSQL_USER", "dev")
	set("MYSQL_PASSWORD", "cdfsdb2023*/")
	set("MYSQL_HOST", "20.38.35.139")
	set("MYSQL_PORT", "3306")
	set("MYSQL_DATABASE", "codefest")
}

// Esta función auxiliar ayuda a crear las variables en el entorno de la aplicación.
func set(s, v string) {
	if err := os.Setenv(s, v); err != nil {
		log.Fatal(err)
	}
}
