package todos

import "time"

type Core struct {
	ID          int
	ActivitysID int
	Title       string
	IsActive    bool
	Priority    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
}

type Data interface {
}
