package service

import "go-testing-example/stub/application/service"

type StubComputeSvc struct {
	service.ComputeSvc
	StartFunc  func() error
	StopFunc   func() error
	RebootFunc func() error
}

func (stub StubComputeSvc) Start() error {
	return stub.StartFunc()
}

func (stub StubComputeSvc) Stop() error {
	return stub.StopFunc()
}

func (stub StubComputeSvc) Reboot() error {
	return stub.RebootFunc()
}
