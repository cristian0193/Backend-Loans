package entity

type Status struct {
	Id          int32  `gorm:"TYPE:SERIAL;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	Description string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:description" json:"description"`
}

func (Status) TableName() string {
	return "Status"
}
