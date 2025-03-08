package repository

import (
	"golang-project/pkg/entity"

	"gorm.io/gorm"
)

type IExampleRepository interface {
	// ExampleMethod
	GetById(id int) (string, error)
	Create(input entity.Example) (id uint, err error)
	Delete(id uint) (err error)
	TestGetById(id uint) (result *entity.Example, err error)
	Get() (result []entity.Example, err error)
	Update(input entity.Example) (err error)
}

type ExampleRepository struct {
	DB *gorm.DB
}

func MakeIExampleRepository(db *gorm.DB) IExampleRepository {
	return &ExampleRepository{
		DB: db,
	}
}

func (r *ExampleRepository) GetById(id int) (string, error) {
	return "test repository", nil
}

func (r ExampleRepository) Create(input entity.Example) (id uint, err error) {
	tx := r.DB.Create(&input)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return input.ID, nil
}

func (r ExampleRepository) Update(input entity.Example) (err error) {
	tx := r.DB.Save(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r ExampleRepository) Delete(id uint) (err error) {
	tx := r.DB.Delete(&entity.Example{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r ExampleRepository) TestGetById(id uint) (result *entity.Example, err error) {
	tx := r.DB.First(&result, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (r ExampleRepository) Get() (result []entity.Example, err error) {
	tx := r.DB.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
