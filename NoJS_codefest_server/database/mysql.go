package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db      *sql.DB
	once    sync.Once
	dbError error
)

// MySQLClient devuelve una instancia única de la conexión a la base de datos MySQL.
func MySQLClient() (*sql.DB, error) {
	once.Do(func() {
		var err error
		db, err = connectMySQL()
		if err != nil {
			dbError = fmt.Errorf("error al conectar a MySQL: %v", err)
			fmt.Println(dbError)
		}
	})
	return db, dbError
}

// connectMySQL establece la conexión con la base de datos MySQL.
func connectMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 2)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("MySQL: Connected!")
	return db, nil
}
