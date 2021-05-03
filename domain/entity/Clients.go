package entity

type Clients struct {
	Identification int32  `gorm:"TYPE:INT8;PRIMARY_KEY;NOT NULL;COLUMN:identification" json:"identification"`
	FullName       string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:fullname" json:"fullname"`
	Address        string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:address" json:"address"`
	Mobile         string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:mobile" json:"mobile"`
}

func (Clients) TableName() string {
	return "Clients"
}
