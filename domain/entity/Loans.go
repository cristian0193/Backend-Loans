package entity

import "time"

type Loans struct {
	Id                   int32     `gorm:"TYPE:SERIAL;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	IdentificationClient int32     `gorm:"TYPE:INT4;NOT NULL;COLUMN:identification_client" json:"identification_client"`
	BorrowedValue        float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:borrowed_value" json:"borrowed_value"`
	InterestPercentage   float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:interest_percentage" json:"interest_percentage"`
	PaidValue            float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:paid_value" json:"paid_value"`
	PendingValue         float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:pending_value" json:"pending_value"`
	InterestPaid         float32   `gorm:"TYPE:NUMERIC;NOT NULL;COLUMN:interest_paid" json:"interest_paid"`
	IdStatus             int       `gorm:"TYPE:INT4;NOT NULL;COLUMN:id_status" json:"id_status"`
	CreationDate         time.Time `gorm:"TYPE:DATE;NOT NULL;COLUMN:creation_date" json:"creation_date"`
	Client               Clients   `gorm:"ForeignKey:identification_client;AssociationForeignKey:identification" json:"client"`
	Interest             Interests `gorm:"ForeignKey:id_loan;AssociationForeignKey:id" json:"interest"`
	State                Status    `gorm:"ForeignKey:id_status;AssociationForeignKey:id" json:"state"`
}

func (Loans) TableName() string {
	return "Loans"
}
