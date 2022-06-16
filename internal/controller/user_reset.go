package controller

import (
	v1 "GaAdmin/api/v1"
	"context"
)

type cUserReset struct{}

var UserReset = cUserReset{}

// 使用手机验证码重置
func (c *cUserReset) ResetMobile(ctx context.Context, req *v1.UserResetMobileReq) (*v1.UserResetMobileRes, error) {
	var (
		res *v1.UserResetMobileRes
	)

	return res, nil
}

// 使用邮件验证码重置
func (c *cUserReset) ResetEmail(ctx context.Context, req *v1.UserResetMobileReq) (*v1.UserResetMobileRes, error) {
	var (
		res *v1.UserResetMobileRes
	)

	return res, nil
}

// 使用安全问答重置
func (c *cUserReset) ResetQuestion(ctx context.Context, req *v1.UserResetQuestionReq) (*v1.UserResetQuestionRes, error) {
	var (
		res *v1.UserResetQuestionRes
	)

	return res, nil
}
