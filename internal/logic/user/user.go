package user

import (
	"GaAdmin/internal/dao"
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/do"
	"GaAdmin/internal/model/entity"
	"GaAdmin/internal/service"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// 创建用户
// - 服务层保证至少一种登录方式
// - 账号|手机号|邮箱其中一个必填
// - 当存在账号，则密码必填
func (s *sUser) CreateUser(ctx context.Context, in *model.UserCreateInput) (*entity.User, error) {
	var (
		data      *do.User
		ent       *entity.User
		available bool
		err       error
	)

	// 使 账号|手机号|邮箱 其中一个必填
	if len(in.Account) == 0 && len(in.Email) == 0 && len(in.Mobile) == 0 {
		return nil, gerror.New("missing passport field")
	}
	// 账号防重，如果有
	if len(in.Account) > 0 {
		if available, err = s.isUserAccountAvailable(ctx, in.Account); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("account is already exists: %s", in.Account)
		}
		if len(in.Password) == 0 {
			return nil, gerror.New("password cannot be empty")
		}
	} else {
		// 随机8位英文，重复由数据库抛出异常
		in.Account = grand.Letters(8)
	}
	// 手机号防重，如果有
	if len(in.Mobile) > 0 {
		if available, err = s.isUserMobileAvailable(ctx, in.Mobile); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("mobile is already exists: %s", in.Mobile)
		}
	}
	// Email防重，如果有
	if len(in.Email) > 0 {
		if available, err = s.isUserEmailAvailable(ctx, in.Email); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("email is already exists: %s", in.Email)
		}
	}
	// 支持无密码创建
	if len(in.Password) == 0 {
		in.Password = grand.Letters(6)
	}

	var (
		salt     = grand.Letters(4)
		insertId int64
	)

	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	data.Uuid = guid.S()
	data.Salt = salt
	data.Password = s.mustEncryptPasword(in.Password, salt)
	// 创建实体
	if err = dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insertId, err = dao.User.Ctx(ctx).Data(data).InsertAndGetId()
		return err
	}); err != nil {
		return nil, err
	}
	ent, _ = s.GetUser(ctx, uint(insertId))

	return ent, nil
}

// 获取用户
func (s *sUser) GetUser(ctx context.Context, userId uint) (*entity.User, error) {
	var (
		ent *entity.User
		err error
	)

	// 扫描数据
	if err = dao.UserGroup.Ctx(ctx).Where(do.User{
		Id: userId,
	}).Scan(&ent); err != nil {
		return nil, err
	}

	return ent, nil
}

// 获取用户集
func (s *sUser) GetUsers(ctx context.Context, userIds []uint) ([]*entity.User, error) {
	var (
		list []*entity.User
		err  error
	)

	// 扫描数据
	if err = dao.UserGroup.Ctx(ctx).WhereIn("id", userIds).Scan(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// 使用uuid获取用户
func (s *sUser) GetUserByUuid(ctx context.Context, uuid string) (*entity.User, error) {
	var (
		ent *entity.User
		err error
	)

	// 扫描数据
	if err = dao.UserGroup.Ctx(ctx).Where(do.User{
		Uuid: uuid,
	}).Scan(&ent); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("user not exist: %d", uuid)
	}

	return ent, nil
}

// 修改用户
func (s *sUser) UpdateUser(ctx context.Context, in *model.UserUpdateInput) (*entity.User, error) {
	var (
		data      *do.User
		ent       *entity.User
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetUser(ctx, in.UserId); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("user not exist: %d", in.UserId)
	}
	// 账户防重，如果有
	if len(in.Account) > 0 {
		if available, err = s.isUserAccountAvailable(ctx, in.Account, []uint{ent.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("account is already exists: %s", in.Account)
		}
	}
	// 手机号防重，如果有
	if len(in.Mobile) > 0 {
		if available, err = s.isUserMobileAvailable(ctx, in.Mobile, []uint{ent.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("mobile is already exists: %s", in.Mobile)
		}
	}
	// 邮箱防重，如果有
	if len(in.Account) > 0 {
		if available, err = s.isUserEmailAvailable(ctx, in.Email, []uint{ent.Id}...); err != nil {
			return nil, err
		}
		if !available {
			return nil, gerror.Newf("email is already exists: %s", in.Email)
		}
	}
	// 转换数据
	if err = gconv.Struct(in, &data); err != nil {
		return nil, err
	}
	// 支持密码为空时不更新
	if len(in.Password) > 0 {
		var salt = grand.Letters(4)
		data.Salt = salt
		data.Password = s.mustEncryptPasword(in.Password, salt)
	} else {
		data.Password = nil
	}
	// 更新实体
	if err = dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.UserId).Data(data).Update()
		return err
	}); err != nil {
		return nil, err
	}
	ent, _ = s.GetUser(ctx, in.UserId)

	return ent, nil
}

// 删除用户
func (s *sUser) DeleteUser(ctx context.Context, id uint) error {
	var (
		ent *entity.User
		err error
	)

	// 扫描数据
	if ent, err = s.GetUser(ctx, id); err != nil {
		return err
	}
	if ent == nil {
		return gerror.Newf("user is not exists")
	}

	// 删除数据
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Delete()
		return err
	})
}

// 获取列表
func (s *sUser) GetUserPage(ctx context.Context, in *model.Page) (*model.UserPageOutput, error) {
	var (
		m    = dao.User.Ctx(ctx)
		out  = &model.UserPageOutput{}
		list []*entity.User
		err  error
	)

	// 分页默认值
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}
	// 组装条件
	if len(in.Condition) > 0 {
		m.Where(in.Condition)
	}
	// 扫描列表
	if err = m.Page(in.Page, in.Size).Order(in.Order).Scan(&list); err != nil {
		return nil, err
	}
	out.List = list
	// 统计分页
	if out.Pager.Total, err = m.Count(); err != nil {
		return nil, err
	}
	out.Size = in.Size
	out.Page = in.Page

	return out, err
}

// 获取当前用户(用于前台)
func (s *sUser) GetCurrentUser(ctx context.Context) (*entity.User, error) {
	var (
		ent *entity.User
		err error
	)

	err = dao.User.Ctx(ctx).Where(do.User{
		Id: service.Context().Get(ctx).User.Id,
	}).Scan(&ent)
	if ent == nil {
		err = gerror.Newf("user not exist: %s", service.Context().Get(ctx).User.Id)
	}
	return ent, err
}

// 修改用户账户(用于前台)
func (s *sUser) UpdateUserAccount(ctx context.Context, account string) error {
	var (
		ent       *entity.User
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetCurrentUser(ctx); err != nil {
		return err
	}
	// 检测防重
	if available, err = s.isUserAccountAvailable(ctx, account, []uint{ent.Id}...); err != nil {
		return err
	}
	if !available {
		return gerror.Newf("account is already exists: %s", account)
	}

	// 更新保存
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).
			Where(dao.User.Columns().Id, ent.Id).
			Data(dao.User.Columns().Account, account).
			Update()
		return err
	})
}

// 修改用户手机号(用于前台)
func (s *sUser) UpdateCurrentUserMobile(ctx context.Context, mobile string) error {
	var (
		ent       *entity.User
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetCurrentUser(ctx); err != nil {
		return err
	}
	// 检测防重
	if available, err = s.isUserMobileAvailable(ctx, mobile, []uint{ent.Id}...); err != nil {
		return err
	}
	if !available {
		return gerror.Newf("mobile is already exists: %s", mobile)
	}

	// 更新保存
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).
			Where(dao.User.Columns().Id, ent.Id).
			Data(dao.User.Columns().Mobile, mobile).
			Update()
		return err
	})
}

// 修改用户邮箱(用于前台)
func (s *sUser) UpdateCurrentUserEmail(ctx context.Context, email string) error {
	var (
		ent       *entity.User
		err       error
		available bool
	)

	// 扫描数据
	if ent, err = s.GetCurrentUser(ctx); err != nil {
		return err
	}
	// 检测防重
	if available, err = s.isUserEmailAvailable(ctx, email, []uint{ent.Id}...); err != nil {
		return err
	}
	if !available {
		return gerror.Newf("email is already exists: %s", email)
	}

	// 更新保存
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).
			Where(dao.User.Columns().Id, ent.Id).
			Data(dao.User.Columns().Email, email).
			Update()
		return err
	})
}

// 修改用户密码(用于前台)
func (s *sUser) UpdateCurrentUserPassword(ctx context.Context, password string) error {
	var (
		ent *entity.User
		err error
	)

	// 扫描数据
	if ent, err = s.GetCurrentUser(ctx); err != nil {
		return err
	}
	password = s.mustEncryptPasword(password, grand.Letters(4))

	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).
			Where(dao.User.Columns().Id, ent.Id).
			Data(dao.User.Columns().Password, password).
			Update()
		return err
	})
}

// 账户|手机号|邮箱 + 密码 登录
func (s *sUser) SignPassport(ctx context.Context, in *model.UserSignPassportInput) (*entity.User, error) {
	var (
		ent      *entity.User
		err      error
		isNoFind bool
	)

	// 优先尝试找手机号
	if err = g.Validator().Rules("phone").Run(ctx); err != nil {
		if isNoFind, err = s.isUserMobileAvailable(ctx, in.Passport); err != nil {
			return nil, err
		}
		if !isNoFind {
			if err = dao.User.Ctx(ctx).Where(do.User{
				Mobile: in.Passport,
			}).Scan(&ent); err != nil {
				return nil, err
			}
		}
	}
	// 手机号找不到，尝试找邮箱
	if ent == nil {
		if err = g.Validator().Rules("email").Run(ctx); err != nil {
			if isNoFind, err = s.isUserEmailAvailable(ctx, in.Passport); err != nil {
				return nil, err
			}
			if !isNoFind {
				if err = dao.User.Ctx(ctx).Where(do.User{
					Email: in.Passport,
				}).Scan(&ent); err != nil {
					return nil, err
				}
			}
		}
	}
	// 邮箱还是找不到，最后找账号
	if ent == nil {
		if isNoFind, err = s.isUserAccountAvailable(ctx, in.Passport); err != nil {
			return nil, err
		}
		if !isNoFind {
			if err = dao.User.Ctx(ctx).Where(do.User{
				Account: in.Passport,
			}).Scan(&ent); err != nil {
				return nil, err
			}
		}
	}
	// 账号密码匹配
	if ent == nil || s.mustEncryptPasword(in.Password, ent.Salt) != ent.Password {
		return nil, gerror.New("passport or password not correct.")
	}

	return ent, nil
}

// 手机号 + 验证码 登录
func (s *sUser) SignMobile(ctx context.Context, in *model.UserSignMobile) (*entity.User, error) {
	var (
		ent *entity.User
		err error
	)

	// 查找用户
	if err = dao.User.Ctx(ctx).Where(dao.User.Columns().Mobile, in.Mobile).Scan(&ent); err != nil {
		return nil, err
	}
	if ent == nil {
		return nil, gerror.Newf("mobile is not find: %s", in.Mobile)
	}
	// 校验验证码
	service.Sms().Verify(ctx, in.Captcha, "signin")

	return ent, nil
}

// 检测账号
func (s *sUser) isUserAccountAvailable(ctx context.Context, account string, notIds ...uint) (bool, error) {
	var (
		m     = dao.User.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.User.Columns().Id, v)
	}
	if count, err = m.Where(dao.User.Columns().Account, account).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}

// 检测手机号
func (s *sUser) isUserMobileAvailable(ctx context.Context, mobile string, notIds ...uint) (bool, error) {
	var (
		m     = dao.User.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.User.Columns().Id, v)
	}
	if count, err = m.Where(dao.User.Columns().Mobile, mobile).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}

// 检测Email
func (s *sUser) isUserEmailAvailable(ctx context.Context, email string, notIds ...uint) (bool, error) {
	var (
		m     = dao.User.Ctx(ctx)
		count int
		err   error
	)

	// 过滤统计
	for _, v := range notIds {
		m = m.WhereNot(dao.User.Columns().Id, v)
	}
	if count, err = m.Where(dao.User.Columns().Email, email).Count(); err != nil {
		return false, err
	}

	return count == 0, nil
}

// 密码盐加密
func (s *sUser) mustEncryptPasword(password, salt string) string {
	return gmd5.MustEncryptString(password + salt)
}
