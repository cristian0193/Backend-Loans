package entity

import "time"

type Payments struct {
	Id          int32     `gorm:"TYPE:SERIAL;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	IdLoan      int32     `gorm:"TYPE:INT4;NOT NULL;COLUMN:id_loan" json:"id_loan"`
	Capital     float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:capital" json:"capital"`
	Interest    float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:interest" json:"interest"`
	Balance     float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:balance" json:"balance"`
	PaymentDate time.Time `gorm:"TYPE:DATE;NOT NULL;COLUMN:payment_date" json:"payment_date"`
	IdType      int       `gorm:"TYPE:INT4;NOT NULL;COLUMN:id_type" json:"id_type"`
	Type        Types     `gorm:"ForeignKey:id_type;AssociationForeignKey:id" json:"type"`
}

func (Payments) TableName() string {
	return "Payments"
}
