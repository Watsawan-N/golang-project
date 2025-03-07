package service

import (
	"golang-project/pkg/errs"
	"golang-project/pkg/helper"
	"golang-project/pkg/repository"
)

type IExampleService interface {
	// ExampleMethod
	GetById(id int) (string, *errs.Error)
}

type ExampleService struct {
	// ExampleService
	Common            helper.ICommon
	ExampleRepository repository.IExampleRepository
	UseCase           IExampleService
}

func MakeIExampleService(common helper.ICommon, repo repository.IExampleRepository) IExampleService {
	return &ExampleService{
		Common:            common,
		ExampleRepository: repo,
	}
}

func (s *ExampleService) GetById(id int) (string, *errs.Error) {
	res, errorRes := s.ExampleRepository.GetById(id)
	if errorRes != nil {
		return res, errs.NewInternalServerError(errorRes.Error())
	}

	return res, nil
}
