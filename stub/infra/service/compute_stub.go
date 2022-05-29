package service

import "go-testing-example/stub/application/service"

type StubComputeSvc struct {
	service.ComputeSvc
	StartFunc  func() error
	StopFunc   func() error
	RebootFunc func() error
}

func (stub StubComputeSvc) Start() error {
	err := stub.StartFunc()
	return err
}

func (stub StubComputeSvc) Stop() error {
	err := stub.StopFunc()
	return err
}

func (stub StubComputeSvc) Reboot() error {
	err := stub.RebootFunc()
	return err
}
