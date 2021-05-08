package entity

type Users struct {
	UserName string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:username" json:"username"`
	Password string `gorm:"TYPE:VARCHAR;NOT NULL;COLUMN:password" json:"password"`
}

func (Users) TableName() string {
	return "Users"
}
