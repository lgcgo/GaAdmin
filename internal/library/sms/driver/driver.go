package driver

// 驱动结构体
type Driver interface {
	// 发送短信
	Send(phone string, message Message, config map[string]string) bool
}

// 短信结构体
type Message struct {
	Template string
	Data     map[string]string
	Content  string
}
