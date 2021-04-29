package persistence

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
	"Backend-Loans/utils"

	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func InitProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}

func (repo *ProductRepositoryImpl) GetByIdMarket(headers dto.Headers) ([]entity.Product, error) {

	var product = []entity.Product{}

	err := repo.db.Where("id_market = ?", headers.IdMarket).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *ProductRepositoryImpl) Insert(productDto dto.ProductDto, headers dto.Headers) error {

	var product = entity.Product{}

	model.Copy(&product, productDto)
	product.IdMarket = headers.IdMarket

	err := repo.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
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
}
