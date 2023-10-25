package models

type Project struct {
	Pr_id     int    `json:"project_id" form:"project_id"`
	Pr_title  string `json:"title" form:"title"`
	Pr_desc   string `json:"desc" form:"desc"`
	Pr_url    string `json:"url_video" form:"url_video"`
	Pr_banner string `json:"url_banner"`
	Pr_status bool   `json:"status"`
}
