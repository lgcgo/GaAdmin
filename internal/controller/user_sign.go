package controller

import (
	v1 "GaAdmin/api/v1"
	"GaAdmin/internal/consts"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type cUserSign struct{}

var UserSign = cUserSign{}

// 员工注册并授权(前台用户注册)
func (c *cUserSign) SignUp(ctx context.Context, req *v1.UserSignUpReq) (*v1.UserSignUpRes, error) {
	var (
		res *v1.UserSignUpRes
		err error
		in  *model.UserCreateInput
		out *model.TokenOutput
		ent *entity.User
	)

	// 校验验证码
	if err = service.Sms().Verify(ctx, req.Captcha, "register"); err != nil {
		return nil, err
	}
	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 创建实体
	if ent, err = service.User().CreateUser(ctx, in); err != nil {
		return nil, err
	}

	// 签发授权
	var (
		auth = service.Auth()
		role string
	)

	if ent.Id == uint(consts.RootAdminId) {
		role = "root"
	} else {
		role = "user"
	}
	if err = auth.InitRbac(); err != nil {
		return nil, err
	}
	if out, err = auth.Authorization(ent.Uuid, role); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 账户|手机号|邮箱 + 密码 登录
func (c *cUserSign) SignPassport(ctx context.Context, req *v1.UserSignPassportReq) (*v1.UserSignPassportRes, error) {
	var (
		res *v1.UserSignPassportRes
		err error
		in  *model.UserSignPassportInput
		out *model.TokenOutput
		ent *entity.User
	)

	// 转换参数
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 账号登录
	if ent, err = service.User().SignPassport(ctx, in); err != nil {
		return nil, err
	}

	// 签发授权(普通用户)
	var (
		auth = service.Auth()
		role string
	)
	if ent.Id == uint(consts.RootAdminId) {
		role = "root"
	} else {
		role = "user"
	}
	if out, err = auth.Authorization(ent.Uuid, role); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 手机号 + 验证码 登录
func (c *cUserSign) SignMobile(ctx context.Context, req *v1.UserSignMobileReq) (*v1.UserSignMobileRes, error) {
	var (
		res *v1.UserSignMobileRes
		err error
		in  *model.UserSignMobile
		out *model.TokenOutput
		ent *entity.User
	)

	// 转换请求
	if err = gconv.Struct(req, &in); err != nil {
		return nil, err
	}
	// 手机号登录
	if ent, err = service.User().SignMobile(ctx, in); err != nil {
		return nil, err
	}

	// 签发授权
	var (
		auth = service.Auth()
		role string
	)

	if ent.Id == uint(consts.RootAdminId) {
		role = "root"
	} else {
		role = "user"
	}
	if out, err = auth.Authorization(ent.Uuid, role); err != nil {
		return nil, err
	}

	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 刷新Token
func (c *cUserSign) Refresh(ctx context.Context, req *v1.UserSignRefreshReq) (*v1.UserSignRefreshRes, error) {
	var (
		auth = service.Auth()
		res  *v1.UserSignRefreshRes
		out  *model.TokenOutput
		err  error
	)

	// 刷新授权
	if out, err = auth.RefreshAuthorization(req.RefreshToken); err != nil {
		return nil, err
	}
	// 转换响应
	if err = gconv.Struct(out, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 注销登录
func (c *cUserSign) SignOut(ctx context.Context, req *v1.UserSignOutReq) (*v1.UserSignOutRes, error) {
	// 这里可以添加token拉黑操作等

	return &v1.UserSignOutRes{}, nil
}
