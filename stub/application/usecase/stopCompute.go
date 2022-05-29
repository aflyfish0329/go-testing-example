package usecase

import (
	"go-testing-example/stub/application/service"
)

func NewStopComputeUseCase(svc service.ComputeSvc) StopComputeUseCase {
	uc := StopComputeUseCase{
		computeSvc: svc,
	}

	return uc
}

type StopComputeUseCase struct {
	computeSvc service.ComputeSvc
}

func (uc StopComputeUseCase) Execute() error {
	// some logic

	// external dependency
	err := uc.computeSvc.Stop()

	// some logic

	return err
}
