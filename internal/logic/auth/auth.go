package auth

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lgcgo/rbac"
)

type sAuth struct {
	r *rbac.Rbac
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

func (s *sAuth) Init() error {
	var (
		ctx                       = gctx.New()
		tokenSignKeyCfg           = g.Cfg().MustGet(ctx, "rbac.tokenSignKey").Bytes()
		tokenIssuerCfg            = g.Cfg().MustGet(ctx, "rbac.tokenIssuer").String()
		accessTokenExpireTimeCfg  = g.Cfg().MustGet(ctx, "rbac.accessTokenExpireTime").String()
		refreshTokenExpireTimeCfg = g.Cfg().MustGet(ctx, "rbac.refreshTokenExpireTime").String()
		policyFilePathCfg         = g.Cfg().MustGet(ctx, "rbac.policyFilePath").String()

		accessTokenExpireTime  time.Duration
		refreshTokenExpireTime time.Duration
		r                      *rbac.Rbac
		err                    error
	)

	if accessTokenExpireTime, err = time.ParseDuration(accessTokenExpireTimeCfg); err != nil {
		return err
	}
	if refreshTokenExpireTime, err = time.ParseDuration(refreshTokenExpireTimeCfg); err != nil {
		return err
	}

	if r, err = rbac.New(rbac.Settings{
		TokenSignKey:           tokenSignKeyCfg,
		TokenIssuer:            tokenIssuerCfg,
		PolicyFilePath:         policyFilePathCfg,
		AccessTokenExpireTime:  accessTokenExpireTime,
		RefreshTokenExpireTime: refreshTokenExpireTime,
	}); err != nil {
		return err
	}
	s.r = r

	return nil
}

// 签发授权
func (s *sAuth) Authorization(subject string, role string) (*model.TokenOutput, error) {
	var (
		err   error
		token *rbac.Token
		out   *model.TokenOutput
	)

	// 初始化
	if err = s.Init(); err != nil {
		return nil, err
	}
	if token, err = s.r.Authorization(subject, role); err != nil {
		return nil, err
	}
	if err = gconv.Struct(out, &token); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *sAuth) RefreshAuthorization(ticket string) (*model.TokenOutput, error) {
	var (
		err   error
		token *rbac.Token
		out   *model.TokenOutput
	)

	// 初始化
	if err = s.Init(); err != nil {
		return nil, err
	}
	if token, err = s.r.RefreshAuthorization(ticket); err != nil {
		return nil, err
	}
	if err = gconv.Struct(out, &token); err != nil {
		return nil, err
	}

	return out, nil
}

// 验证Token
func (s *sAuth) VerifyToken(ticket string) (g.Map, error) {
	var (
		out g.Map
		err error
	)

	// 初始化
	if err = s.Init(); err != nil {
		return nil, err
	}

	if out, err = s.r.VerifyToken(ticket); err != nil {
		return nil, err
	}

	return out, nil
}

// 验证路由
func (s *sAuth) VerifyRequest(path, method, role string) error {
	var (
		err error
	)

	// 初始化
	if err = s.Init(); err != nil {
		return err
	}

	return s.r.VerifyRequest(path, method, role)
}
