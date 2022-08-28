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
	InsertData(insert Core) (data Core, row int, err error)
}

type Data interface {
	InsertData(insert Core) (data Core, row int, err error)
}
