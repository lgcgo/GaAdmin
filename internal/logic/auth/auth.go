package auth

import "GaAdmin/internal/service"

type sAuth struct{}

func Init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}
