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
	GetData(id int) (data Core, row int, err error)
	InsertData(insert Core) (data Core, row int, err error)
	UpdateData(id int, insert Core) (data Core, row int, err error)
}

type Data interface {
	GetData(id int) (data Core, row int, err error)
	InsertData(insert Core) (data Core, row int, err error)
	UpdateData(id int, insert Core) (data Core, row int, err error)
	UniqueData(insert Core) (row int, err error)
}
