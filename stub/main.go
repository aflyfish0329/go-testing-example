package main

import (
	"go-testing-example/stub/infra/service"

	"go-testing-example/stub/application/usecase"
)

func main() {
	uc := usecase.NewStopComputeUseCase(service.ComputeSvc{})
	uc.Execute()
}
