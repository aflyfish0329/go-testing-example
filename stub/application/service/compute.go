package service

type ComputeSvc interface {
	Start() error
	Stop() error
	Reboot() error
}
