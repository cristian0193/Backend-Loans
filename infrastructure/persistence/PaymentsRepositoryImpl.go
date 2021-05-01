package persistence

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
)

type PaymentsRepositoryImpl struct {
	db *gorm.DB
}

func InitPaymentsRepositoryImpl(db *gorm.DB) *PaymentsRepositoryImpl {
	return &PaymentsRepositoryImpl{db}
}

func (repo *PaymentsRepositoryImpl) Insert(paymentDto dto.PaymentDto) error {

	var payment = entity.Payments{}

	model.Copy(&payment, paymentDto)
	payment.PaymentDate = time.Now()

	err := repo.db.Create(&payment).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PaymentsRepositoryImpl) UpdateBalance(id int32, balance float32) error {

	err := repo.db.Table("Payments").Where("id = ?", id).
		Updates(map[string]interface{}{"balance": balance}).Error

	if err != nil {
		return err
	}
	return nil
}

func (repo *PaymentsRepositoryImpl) FindLastBalance(idLoan int32) (dto.LastBalanceDto, error) {

	var lastBalance = dto.LastBalanceDto{}
	var query = `SELECT p.id, p.balance, l.interest_percentage 
					FROM "Loans" l
					inner join "Payments" p on l.id = p.id_loan
					WHERE p.id_loan = ?
					ORDER BY id 
					DESC LIMIT 1;`

	err := repo.db.Raw(query, idLoan).Scan(&lastBalance).Error
	if err != nil {
		return lastBalance, err
	}
	return lastBalance, nil
}
