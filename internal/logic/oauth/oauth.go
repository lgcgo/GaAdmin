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
func (s *sOauth) Authorization(ctx context.Context, subject string, issueRoles []string) (*model.TokenOutput, error) {
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

	iClaims = &jwt.IssueClaims{
		Subject:    subject,
		IssueRoles: issueRoles,
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

	// 组装数据
	expiresIn = currentTime.Add(atExpireTime).Sub(currentTime).Seconds()
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
	if claims, err = jwt.NewJwt().ParseToken(ticket); err != nil {
		return nil, err
	}
	// 非法动作签名
	if claims["act"] != "renew" {
		return nil, gerror.New("claims act not correct.")
	}
	return s.Authorization(ctx, claims["sub"].(string), claims["isr"].([]string))
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
	if claims["act"] != "issue" {
		return nil, gerror.New("claims act not correct.")
	}

	return claims, nil
}

// 检查路径
func (s *sOauth) CheckPath(r *ghttp.Request, issueRoles []string) (bool, error) {
	var (
		plicy *casbin.Policy
		err   error
		ok    bool
	)
	for _, issueRole := range issueRoles {
		plicy = &casbin.Policy{
			Subject: issueRole,
			Object:  r.URL.Path,
			Action:  r.Method,
		}
		if ok, err = casbin.NewCasbin().Verify(plicy); err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, gerror.New("no permission")
}

// 更新授权政策
func (s *sOauth) SavePolicy(ctx context.Context) error {
	var (
		policys []*casbin.Policy
		roles   []*casbin.Role
		ugList  []*entity.UserGroup
		ugaList []*entity.UserGroupAccess
		err     error
	)

	// 组装角色政策
	if ugList, err = service.User().GetAllGroup(ctx); err != nil {
		return err
	}
	for _, v := range ugList {
		if v.ParentId == 0 {
			roles = append(roles, &casbin.Role{
				ParentSubject: "root",
				Subject:       v.Name,
			})
		} else {
			var parentName string
			if parentName, err = service.User().GetGroupName(ctx, v.ParentId); err != nil {
				return err
			}
			roles = append(roles, &casbin.Role{
				ParentSubject: parentName,
				Subject:       v.Name,
			})
		}
	}
	// 组装节点政策
	if ugaList, err = service.User().GetAllGroupAccess(ctx); err != nil {
		return err
	}
	for _, v := range ugaList {
		var (
			name string
			rule *entity.AuthRule
		)
		if name, err = service.User().GetGroupName(ctx, v.GroupId); err != nil {
			return err
		}
		rule, err = service.Auth().GetRule(ctx, v.AuthRuleId)
		policys = append(policys, &casbin.Policy{
			Subject: name,
			Object:  rule.Path,
			Action:  rule.Method,
		})
	}
	// 保存配置文件
	if err = casbin.NewCasbin().SavePolicyCsv(policys, roles); err != nil {
		return err
	}

	return nil
}
