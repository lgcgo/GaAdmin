package sms

import (
	"GaAdmin/internal/library/sms/driver"
	"sync"
)

// 操作类
type SMS struct {
	Driver driver.Driver
}

// 单例模式
var once sync.Once

// 实例声明
var insSMS *SMS

// 获取阿里云单例
func NewAliyunSMS() *SMS {
	once.Do(func() {
		insSMS = &SMS{
			Driver: &driver.Aliyun{},
		}
	})
	return insSMS
}

// 发送
func (s *SMS) Send(phone string, message driver.Message) bool {
	config := map[string]string{
		"test": "abc",
	}
	return s.Driver.Send(phone, message, config)
}
