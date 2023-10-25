package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Us_id       int    `json:"user_id"`
	Us_name     string `json:"name" form:"name"`
	Us_email    string `json:"email" form:"email"`
	Us_pwd      string `json:"password,omitempty" form:"pwd"`
	Us_avatar   string `json:"avatar_image"`
	Us_faculty  string `json:"faculty" form:"faculty"`
	Us_semester int    `json:"semester" form:"semester"`
	Us_status   bool   `json:"status"`
	Ty_id       int    `json:"user_type"`
	Token       string `json:"token"`
	jwt.RegisteredClaims
}
