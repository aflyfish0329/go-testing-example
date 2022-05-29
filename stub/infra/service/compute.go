package service

import "fmt"

type ComputeSvc struct{}

func (svc ComputeSvc) Start() error {
	fmt.Println("cloud start")
	return nil
}

func (svc ComputeSvc) Stop() error {
	fmt.Println("cloud stop")
	return nil
}

func (svc ComputeSvc) Reboot() error {
	fmt.Println("cloud reboot")
	return nil
}
