package activitys

import "time"

type Core struct {
	ID        int
	Email     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
}

type Data interface {
}
