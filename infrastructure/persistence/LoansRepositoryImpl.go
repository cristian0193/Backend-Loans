package persistence

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
)

type LoansRepositoryImpl struct {
	db *gorm.DB
}

func InitLoansRepositoryImpl(db *gorm.DB) *LoansRepositoryImpl {
	return &LoansRepositoryImpl{db}
}

func (repo *LoansRepositoryImpl) Insert(loansDto dto.LoansDto) (int32, error) {

	layOut := "2006-01-02"
	var loans = entity.Loans{}

	dateStamp, _ := time.Parse(layOut, loansDto.CreationDate)

	model.Copy(&loans, loansDto)
	loans.IdStatus = 2
	loans.CreationDate = dateStamp

	err := repo.db.Create(&loans).Error
	if err != nil {
		return 0, err
	}
	return loans.Id, nil
}

func (repo *LoansRepositoryImpl) FindUserByUser(identification int32) (bool, error) {

	var loans = []entity.Loans{}

	err := repo.db.Where("identification_client = ? and ((id_status = 2) or (id_status = 3))", identification).
		Find(&loans).Error

	if err != nil {
		return false, err
	}

	if len(loans) > 0 {
		return true, err
	}

	return false, nil
}

func (repo *LoansRepositoryImpl) FindUserById(id int32) (entity.Loans, error) {

	var loans = entity.Loans{}

	err := repo.db.Where("id = ?", id).Find(&loans).Error
	if gorm.IsRecordNotFoundError(err) {
		return loans, nil
	}

	if err != nil {
		return loans, err
	}
	return loans, nil
}

func (repo *LoansRepositoryImpl) UpdateCalculateById(id int32) error {

	var query = `UPDATE public."Loans"
	SET  paid_value=(SELECT sum(p.capital) 
					 FROM "Payments" p 
					 WHERE p.id_loan = ?), 
		 pending_value=(SELECT ((l.borrowed_value) - sum(p.capital)) as pending
						FROM "Loans" l
						INNER JOIN "Payments" p on l.id = p.id_loan
						WHERE p.id_loan = ?
						GROUP BY l.id), 
		 interest_paid=(SELECT sum(p.interest) 
						 FROM "Payments" p 
						 WHERE p.id_loan = ?), 
		 id_status=(SELECT 
						CASE WHEN l.borrowed_value = sum(p.capital) THEN 1
							 ELSE 2
						END
					FROM "Loans" l
					INNER JOIN "Payments" p on l.id = p.id_loan
					WHERE p.id_loan = ?
					GROUP BY l.id)
	WHERE id=?;`

	err := repo.db.Exec(query, id, id, id, id, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *LoansRepositoryImpl) FindAllLoan(query dto.QueryParameters) ([]entity.Loans, error) {
	var loans = []entity.Loans{}
	var err error

	limit := 6
	pages := int(query.Pages-1) * limit

	if query.Fullname != "" {
		err = repo.db.Joins(`inner join "Clients" on "Clients".identification = "Loans".identification_client`).
			Where(`"Loans".id_status = ? and "Clients".fullname like ?`, query.Status, ("%" + query.Fullname + "%")).
			Preload("Client").Preload("State").
			Order("id DESC").
			Offset(0).
			Limit(limit).
			Find(&loans).Error
	} else {
		err = repo.db.Where(`id_status = ?`, query.Status).
			Preload("Client").Preload("State").
			Order("id DESC").
			Offset(pages).
			Limit(limit).
			Find(&loans).Error
	}

	if gorm.IsRecordNotFoundError(err) {
		return loans, nil
	}

	if err != nil {
		return loans, err
	}
	return loans, nil
}

func (repo *LoansRepositoryImpl) CountAllLoans(status uint) (int, error) {

	var countLoan int
	query := `select count(*) as countLoan from "Loans" where id_status = ?`

	err := repo.db.Raw(query, status).Count(&countLoan).Error

	if gorm.IsRecordNotFoundError(err) {
		return countLoan, nil
	}

	if err != nil {
		return countLoan, err
	}
	return countLoan, nil
}

func (repo *LoansRepositoryImpl) FindInformacionByIdLoan(idLoan int32) (entity.Loans, error) {

	var loans = entity.Loans{}

	err := repo.db.Where("id = ?", idLoan).
		Preload("Client").
		Preload("State").
		Preload("Interest", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = 'ACT'")
		}).
		Find(&loans).Error
	if gorm.IsRecordNotFoundError(err) {
		return loans, nil
	}

	if err != nil {
		return loans, err
	}
	return loans, nil
}
