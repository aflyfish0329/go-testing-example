package usecase

import "go-testing-example/stub/application/service"

type StubComputeSvc struct {
	// 嵌入 interface, 可以不用實作所有的 interface method, 但呼叫時會出錯.
	// 如實作了全部的 method, 則無需內嵌.
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
