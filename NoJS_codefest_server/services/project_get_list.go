package services

import (
	"NoJS_codefest_server/database"
	"NoJS_codefest_server/models"
)

func ProjectGetListService() ([]models.Project, error) {
	db, _ := database.MySQLClient()
	query := "SELECT * FROM projects"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.Pr_id, &project.Pr_title, &project.Pr_desc, &project.Pr_url, &project.Pr_banner, &project.Pr_status)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
