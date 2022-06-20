package org

import (
	"GaAdmin/internal/service"
)

type sOrg struct{}

func init() {
	service.RegisterOrg(New())
}

func New() *sOrg {
	return &sOrg{}
}
