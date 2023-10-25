package services

import (
	"NoJS_codefest_server/database"
	"NoJS_codefest_server/functions"
	"NoJS_codefest_server/models"
	"log"
)

func UserRegisterService(user models.User) error {
	db, err := database.MySQLClient()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (us_name, us_email, us_pwd, us_avatar, us_faculty, us_semester, us_status, ty_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Us_name, user.Us_email, functions.Sha512hash(user.Us_pwd), user.Us_avatar, user.Us_faculty, user.Us_semester, user.Us_status, user.Ty_id)
	if err != nil {
		return err
	}
	return nil
}
