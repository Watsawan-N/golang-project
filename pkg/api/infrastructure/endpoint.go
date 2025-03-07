package infrastructure

import "golang-project/pkg/endpoint"

type Endpoint struct {
	// ExampleEndpoint
	ExampleEndpoint *endpoint.ExampleEndpoint
}

func CreateEndpoint(service Service, helper Helper) Endpoint {
	return Endpoint{
		ExampleEndpoint: endpoint.MakeExampleEndpoint(&service.IExampleService, &helper.ICommon),
	}
}
