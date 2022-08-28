package request

import "testcode/features/activitys"

type Activities struct {
	Email string `json:"email" form:"email"`
	Title string `json:"title" form:"title"`
}

func ToCore(req Activities) activitys.Core {
	return activitys.Core{
		Email: req.Email,
		Title: req.Title,
	}
}
