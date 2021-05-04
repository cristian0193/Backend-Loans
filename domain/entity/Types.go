package entity

type Types struct {
	Id   int32  `gorm:"TYPE:SERIAL;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	Name string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:name" json:"name"`
}

func (Types) TableName() string {
	return "Types"
}
