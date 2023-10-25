package services

import (
	"NoJS_codefest_server/database"
	"NoJS_codefest_server/functions"
	"NoJS_codefest_server/models"
	"log"
)

func UserLoginService(email string, password string) (models.User, error) {
	db, err := database.MySQLClient()
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	err = db.QueryRow("SELECT * FROM users WHERE us_email = ? AND us_pwd = ?", email, functions.Sha512hash(password)).Scan(&user.Us_id, &user.Us_name, &user.Us_email, &user.Us_pwd, &user.Us_avatar, &user.Us_faculty, &user.Us_semester, &user.Us_status, &user.Ty_id)
	if err != nil {
		return models.User{}, err
	}
	user.Us_pwd = ""
	return user, nil
}
