package products

import (
	"errors"
	"golang-adv/4-order-api/pkg/db"
	"log"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Delete(id uint) error {
	product, err := repo.GetById(id)
	if err != nil {
		return err
	}
	result := repo.Database.Delete(product, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return &Product{}, errors.New("Product not found")
	}
	return product, nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	product := Product{}
	result := repo.Database.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	log.Println("Start")
	result := repo.Database.Find(&products)
	log.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
