package model

// 授权返回
type TokenOutput struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    uint
	RefreshToken string
}
