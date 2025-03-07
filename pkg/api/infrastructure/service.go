package infrastructure

import (
	"golang-project/pkg/service"
)

type Service struct {
	IExampleService  service.IExampleService
	IErrorLogService service.IErrorLogService
}

func CreateService(repo Repository, helper Helper) Service {
	return Service{
		IExampleService:  service.MakeIExampleService(helper.ICommon, repo.IExampleRepository),
		IErrorLogService: service.MakeIErrorLogService(repo.IErrorLogRepository),
	}
}
