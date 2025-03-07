package repository

import "gorm.io/gorm"

type IExampleRepository interface {
	// ExampleMethod
	GetById(id int) (string, error)
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
