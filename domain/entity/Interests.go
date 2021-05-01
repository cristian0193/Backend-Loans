package entity

type Interests struct {
	Id     int32   `gorm:"TYPE:SERIAL;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	IdLoan int32   `gorm:"TYPE:INT4;NOT NULL;COLUMN:id_loan" json:"id_loan"`
	Share  float32 `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:share" json:"share"`
	Status string  `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:status" json:"status"`
}

func (Interests) TableName() string {
	return "Interests"
}
