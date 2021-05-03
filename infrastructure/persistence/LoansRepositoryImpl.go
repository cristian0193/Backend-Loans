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

	var loans = entity.Loans{}

	model.Copy(&loans, loansDto)
	loans.IdStatus = 2
	loans.CreationDate = time.Now()

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

func (repo *LoansRepositoryImpl) FindAllLoans() ([]entity.Loans, error) {
	var loans = []entity.Loans{}

	err := repo.db.Preload("Client").Find(&loans).Error
	if gorm.IsRecordNotFoundError(err) {
		return loans, nil
	}

	if err != nil {
		return loans, err
	}
	return loans, nil
}

/* func (repo *LoansRepositoryImpl) GetByIdMarket(headers dto.Headers) ([]entity.Product, error) {

	var product = []entity.Product{}

	err := repo.db.Where("id_market = ?", headers.IdMarket).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *ProductRepositoryImpl) Update(productDto dto.ProductDto, headers dto.Headers) error {

	var product = entity.Product{}

	model.Copy(&product, productDto)
	product.IdMarket = headers.IdMarket

	err := repo.db.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) Delete(idProduct string, headers dto.Headers) error {

	var product = entity.Product{}

	err := repo.db.Where("id = ? and id_market = ?", idProduct, headers.IdMarket).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) GetByIdProduct(idProduct string, headers dto.Headers) (entity.Product, error) {

	var product = entity.Product{}

	err := repo.db.Where("id = ? and id_market = ?", idProduct, headers.IdMarket).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *ProductRepositoryImpl) GetByQueryParameters(queryParameters dto.QueryParameters, headers dto.Headers) ([]entity.Product, error) {
	var product = []entity.Product{}
	queryFilterParameters := &queryParameters

	var queryParameter = utils.FilterQueryParameters(queryFilterParameters)

	var query = `SELECT id, name, price, id_category FROM public."Product" WHERE id_market = ? and ` + queryParameter

	err := repo.db.Raw(query, headers.IdMarket).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
} */
