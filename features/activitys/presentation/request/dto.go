package request

import "testcode/features/activitys"

type Activitys struct {
	Email string `json:"email" form:"email"`
	Title string `json:"title" form:"title"`
}

func ToCore(req Activitys) activitys.Core {
	return activitys.Core{
		Email: req.Email,
		Title: req.Title,
	}
}
