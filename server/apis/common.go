package apis

import "gametools/server/app"

type Query struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (q *Query) GetPage() int {
	if q.Page <= 0 {
		q.Page = 1
	}
	return q.Page
}

func (q *Query) GetLimit() int {
	if q.Limit <= 0 {
		q.Limit = app.GetApp().Config.Server.DefaultPageLimit
	}
	return q.Limit
}

func (q *Query) GetOffset() int {
	page := q.GetPage()
	limit := q.GetLimit()
	return (page - 1) * limit
}
