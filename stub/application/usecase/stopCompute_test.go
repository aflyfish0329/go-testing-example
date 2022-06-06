package usecase_test

import (
	"errors"
	"testing"

	"go-testing-example/stub/application/usecase"
	"go-testing-example/stub/infra/service"

	"github.com/stretchr/testify/assert"
)

func Test_StopComputeUseCase_Succeed(t *testing.T) {
	stub := service.StubComputeSvc{
		StopFunc: func() error {
			return nil
		},
	}

	uc := usecase.NewStopComputeUseCase(stub)
	err := uc.Execute()
	assert.NoError(t, err)
}

func Test_StopComputeUseCase_Failed(t *testing.T) {
	stub := service.StubComputeSvc{
		StopFunc: func() error {
			return errors.New("error")
		},
	}

	uc := usecase.NewStopComputeUseCase(stub)
	err := uc.Execute()
	assert.Error(t, err)
}
