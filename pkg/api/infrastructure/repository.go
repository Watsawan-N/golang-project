package infrastructure

import (
	"golang-project/pkg/repository"

	"gorm.io/gorm"
)

type Repository struct {
	IExampleRepository  repository.IExampleRepository
	IErrorLogRepository repository.IErrorLogRepository
}

func CreateRepository(db *gorm.DB) Repository {
	return Repository{
		IExampleRepository:  repository.MakeIExampleRepository(db),
		IErrorLogRepository: repository.MakeIErrorLogRepository(db),
	}
}
