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

	layOut := "2006-01-02"
	var payment = entity.Payments{}

	dateStamp, _ := time.Parse(layOut, paymentDto.PaymentDate)

	model.Copy(&payment, paymentDto)
	payment.PaymentDate = dateStamp

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

func (repo *PaymentsRepositoryImpl) FindByIdLoan(idLoan int32) ([]entity.Payments, error) {

	var payment = []entity.Payments{}

	err := repo.db.Where("id_loan = ?", idLoan).Preload("Type").Find(&payment).Error
	if err != nil {
		return payment, err
	}
	return payment, nil
}
