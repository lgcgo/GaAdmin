package jwt

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	pkg "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	// 上下文
	Ctx context.Context
	// 加密秘钥
	SignKey []byte
	// 发布者
	Issuer string
}

// StandardClaims 结构体实现了 Claims 接口继承了 Valid() 方法
// JWT 规定了7个官方字段，提供使用:
// - iss (issuer)：发布者
// - sub (subject)：主题
// - iat (Issued At)：生成签名的时间
// - exp (expiration time)：签名过期时间
// - aud (audience)：观众，相当于接受者
// - nbf (Not Before)：生效时间
// - jti (JWT ID)：编号
type Claims struct {
	IssueType string `json:"ist"` // 签发行为, issue=签发,renew=刷新
	IssueRole string `json:"isr"` // 签发角色, 签发的角色名称（允许多角色）
	pkg.RegisteredClaims
}

type IssueClaims struct {
	IssueType string `v:"required"`
	IssueRole string `v:"required"`
	Subject   string `v:"required"`
}

// 实例声明
var insJWT *JWT

func NewJwt() *JWT {
	var (
		ctx = context.Background()
	)

	insJWT = &JWT{
		Ctx:     ctx,
		SignKey: g.Cfg().MustGet(ctx, "jwt.signKey").Bytes(),
		Issuer:  g.Cfg().MustGet(ctx, "jwt.claims.issuer").String(),
	}

	return insJWT
}

// 签发Token
func (j *JWT) IssueToken(iClaims *IssueClaims, expireTime time.Duration) (string, error) {
	var (
		token  *pkg.Token
		ticket string
		err    error
	)

	// 验证字段
	if err = g.Validator().Data(iClaims).Run(j.Ctx); err != nil {
		return "", err
	}
	// 创建签名
	claims := &Claims{
		iClaims.IssueType,
		iClaims.IssueRole,
		pkg.RegisteredClaims{
			Issuer:    j.Issuer,
			Subject:   iClaims.Subject,
			Audience:  []string{},
			ExpiresAt: pkg.NewNumericDate(time.Now().Add(expireTime)),
			NotBefore: pkg.NewNumericDate(time.Now()),
			IssuedAt:  pkg.NewNumericDate(time.Now()),
		},
	}
	// 生成token
	token = pkg.NewWithClaims(pkg.SigningMethodHS256, claims)
	if ticket, err = token.SignedString(j.SignKey); err != nil {
		return "", err
	}

	return ticket, nil
}

// 解析Token
func (j *JWT) ParseToken(ticket string) (map[string]interface{}, error) {
	var (
		token   *pkg.Token
		mClaims map[string]interface{}
		err     error
		ok      bool
	)

	// 解析Token对象
	if token, err = pkg.Parse(ticket, func(token *pkg.Token) (interface{}, error) {
		if _, ok = token.Method.(*pkg.SigningMethodHMAC); !ok {
			return nil, gerror.Newf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SignKey, nil
	}); err != nil {
		return nil, err
	}

	// 验证签名
	if mClaims, ok = token.Claims.(pkg.MapClaims); !ok || !token.Valid {
		return nil, gerror.New("token parse error")
	}

	return mClaims, nil
}
