package todos

import "time"

type Core struct {
	ID              int
	ActivityGroupID int
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Business interface {
	GetAllData(param string) (data []Core, row int, err error)
	GetData(id int) (data Core, row int, err error)
	InsertData(insert Core) (data Core, row int, err error)
	UpdateData(id int, insert Core) (data Core, row int, err error)
	DeleteData(id int) (row int, err error)
}

type Data interface {
	GetAllData(param string) (data []Core, row int, err error)
	GetData(id int) (data Core, row int, err error)
	InsertData(insert Core) (data Core, row int, err error)
	UpdateData(id int, insert Core) (data Core, row int, err error)
	DeleteData(id int) (row int, err error)
}
