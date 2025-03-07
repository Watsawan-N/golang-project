package service

import (
	"golang-project/pkg/entity"
	"golang-project/pkg/errs"
	"golang-project/pkg/model"
	"golang-project/pkg/repository"
)

type IErrorLogService interface {
	Create(input model.ErrorLog) *errs.Error
	DeleteOverDueDate(day int) *errs.Error
}

type ErrorLogService struct {
	ErrorLogRepository repository.IErrorLogRepository
}

func MakeIErrorLogService(errorLogRepository repository.IErrorLogRepository) IErrorLogService {
	return &ErrorLogService{
		ErrorLogRepository: errorLogRepository,
	}
}

func (s ErrorLogService) Create(input model.ErrorLog) *errs.Error {
	inputEntity := entity.ErrorLog{
		Method:         input.Method,
		RequestURI:     input.RequestURI,
		RequestBody:    input.RequestBody,
		RequestTime:    input.RequestTime,
		ResponseTime:   input.ResponseTime,
		ResponseStatus: input.ResponseStatus,
		ResponseBody:   input.ResponseBody,
	}

	err := s.ErrorLogRepository.Create(inputEntity)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (s ErrorLogService) DeleteOverDueDate(day int) *errs.Error {
	err := s.ErrorLogRepository.DeleteOverDueDate(day)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
