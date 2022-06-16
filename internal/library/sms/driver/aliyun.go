package driver

type Aliyun struct{}

func (d *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	return false
}
