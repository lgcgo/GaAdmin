package sms

import (
	"GaAdmin/internal/service"
	"context"
)

type sSms struct{}

func Init() {
	service.RegisterSms(New())
}

func New() *sSms {
	return &sSms{}
}

func (s *sSms) Verify(ctx context.Context, code string, scene string) error {
	return nil
}
