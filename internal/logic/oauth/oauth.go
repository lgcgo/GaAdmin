package oauth

import (
	"GaAdmin/internal/library/casbin"
	"GaAdmin/internal/library/jwt"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	service.RegisterOauth(New())
}

type sOauth struct{}

/*
* 当前实现Oauth的密码授权模式，满足RBAC权限认证需求
 */
func New() *sOauth {
	return &sOauth{}
}

// 签发授权
func (s *sOauth) Authorization(ctx context.Context, subject string, issueRole string) (*model.TokenOutput, error) {
	var (
		pkg             = jwt.NewJwt()
		atExpireTimeCfg = g.Cfg().MustGet(ctx, "jwt.accessToken.expireTime").String()
		rtExpireTimeCfg = g.Cfg().MustGet(ctx, "jwt.refreshToken.expireTime").String()
		currentTime     = time.Now()
		rtExpireTime    time.Duration
		atExpireTime    time.Duration
		res             *model.TokenOutput
		err             error
		accessToken     string
		refreshToken    string
		expiresIn       float64
		iClaims         *jwt.IssueClaims
	)

	// 实例化签名
	iClaims = &jwt.IssueClaims{
		Subject:   subject,
		IssueRole: issueRole,
	}
	// 制作 accessToken
	if atExpireTime, err = time.ParseDuration(atExpireTimeCfg); err != nil {
		return nil, err
	}
	iClaims.IssueType = "issue"
	if accessToken, err = pkg.IssueToken(iClaims, atExpireTime); err != nil {
		return nil, err
	}
	// 制作 refreshToken
	if rtExpireTime, err = time.ParseDuration(rtExpireTimeCfg); err != nil {
		return nil, err
	}
	iClaims.IssueType = "renew"
	if refreshToken, err = pkg.IssueToken(iClaims, rtExpireTime); err != nil {
		return nil, err
	}
	// 获取过期秒数
	expiresIn = currentTime.Add(atExpireTime).Sub(currentTime).Seconds()
	// 组装返回数据
	res = &model.TokenOutput{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
		ExpiresIn:    uint(expiresIn),
	}

	return res, nil
}

// 刷新授权
func (s *sOauth) RefreshAuthorization(ctx context.Context, ticket string) (*model.TokenOutput, error) {
	var (
		claims map[string]interface{}
		err    error
	)

	// 解析token
	if claims, err = jwt.NewJwt().ParseToken(ticket); err != nil {
		return nil, err
	}
	// 校验签发类型
	if claims["ist"] != "renew" {
		return nil, gerror.New("claims ist not correct.")
	}

	return s.Authorization(ctx, claims["sub"].(string), claims["isr"].(string))
}

// 验证授权
func (s *sOauth) ValidAuthorization(r *ghttp.Request) (g.Map, error) {
	var (
		claims g.Map
		err    error
		strArr []string
	)

	// 支持Header传值方式
	if r.Header.Get("Authorization") == "" {
		return nil, gerror.New("headers authorization not exists.")
	}
	strArr = strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	// 支持Bearer方案
	if strArr[0] != "Bearer" {
		return nil, gerror.New("authorization scheme not support.")
	}
	// 解析token票据
	if claims, err = jwt.NewJwt().ParseToken(strArr[1]); err != nil {
		return nil, err
	}
	// 非法动作签名
	if claims["ist"] != "issue" {
		return nil, gerror.New("claims ist not correct.")
	}

	return claims, nil
}

// 检查路径
func (s *sOauth) CheckPath(r *ghttp.Request, issueRole string) (bool, error) {
	var (
		plicy *casbin.Policy
		err   error
		ok    bool
	)

	// 组装政策数据
	plicy = &casbin.Policy{
		Subject: issueRole,
		Object:  r.URL.Path,
		Action:  r.Method,
	}
	// 验证政策权限
	if ok, err = casbin.NewCasbin().Verify(plicy); err != nil {
		return false, err
	}
	if !ok {
		return false, gerror.New("no permission")
	}

	return true, nil
}

// 更新授权政策
func (s *sOauth) SavePolicy(ctx context.Context) error {
	var (
		policys    []*casbin.Policy
		roles      []*casbin.Role
		accessList []*entity.AuthRoleAccess
		roleList   []*entity.AuthRole
		err        error
	)

	// 组装角色权限政策
	if accessList, err = service.Auth().GetAllRoleAccess(ctx); err != nil {
		return err
	}
	for _, v := range accessList {
		var (
			role *entity.AuthRole
			rule *entity.AuthRule
		)

		if role, err = service.Auth().GetRole(ctx, v.RoleId); err != nil {
			return err
		}
		if rule, err = service.Auth().GetRule(ctx, v.RuleId); err != nil {
			return err
		}
		policys = append(policys, &casbin.Policy{
			Subject: role.Name,
			Object:  rule.Path,
			Action:  rule.Method,
		})
	}

	// 组装角色层级政策
	if roleList, err = service.Auth().GetAllRole(ctx); err != nil {
		return err
	}
	for _, v := range roleList {
		if v.ParentId == 0 {
			roles = append(roles, &casbin.Role{
				ParentSubject: "root",
				Subject:       v.Name,
			})
		} else {
			var parentName string
			if parentName, err = service.Auth().GetRoleName(ctx, v.ParentId); err != nil {
				return err
			}
			roles = append(roles, &casbin.Role{
				ParentSubject: parentName,
				Subject:       v.Name,
			})
		}
	}
	// 保存配置文件
	if err = casbin.NewCasbin().SavePolicyCsv(policys, roles); err != nil {
		return err
	}

	return nil
}
