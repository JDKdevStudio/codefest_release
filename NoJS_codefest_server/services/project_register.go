package services

import (
	"NoJS_codefest_server/database"
	"NoJS_codefest_server/models"
	"log"
)

func ProjectRegisterService(project models.Project, user_id int) error {
	db, err := database.MySQLClient()
	if err != nil {
		log.Fatal(err)
	}
	query := "INSERT INTO projects (pr_title, pr_desc, pr_url, pr_banner, pr_status) VALUES (?, ?, ?, ?, ?)"
	response, err := db.Exec(query, project.Pr_title, project.Pr_desc, project.Pr_url, project.Pr_banner, project.Pr_status)
	if err != nil {
		return err
	}
	last_id, err := response.LastInsertId()
	if err != nil {
		return err
	}
	go func() {
		db.Exec("INSERT INTO users_has_projects (us_id, pr_id, uspr_member_type) VALUES(?, ?, ?)", user_id, last_id, 1)
	}()
	return nil
}
