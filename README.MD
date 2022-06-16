# 介绍
GaAdmin是一款基于GoFrame+Antd的高颜值开发框架。

# 主要特性
- 基于 `Casbin` + `Jwt` 高效的 RBAC 权限认证模块
- 统一通用的会员模块，支持多应用下的单点登录
- 待完善...

# 安装使用
- 待完善...

# 前端项目
- 待完善...

# 界面截图
- 待完善...

# 接口清单
```shell
ADDRESS | METHOD |        ROUTE        |                             HANDLER
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | DELETE | /auth/menu          | GaAdmin/internal/controller.(*cAuthMenu).Delete
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /auth/menu          | GaAdmin/internal/controller.(*cAuthMenu).Get
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /auth/menu          | GaAdmin/internal/controller.(*cAuthMenu).Create
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /auth/menu          | GaAdmin/internal/controller.(*cAuthMenu).Update
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /auth/menu-tree     | GaAdmin/internal/controller.(*cAuthMenu).Tree
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | DELETE | /auth/rule          | GaAdmin/internal/controller.(*cAuthRule).Delete
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /auth/rule          | GaAdmin/internal/controller.(*cAuthRule).Get
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /auth/rule          | GaAdmin/internal/controller.(*cAuthRule).Create
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /auth/rule          | GaAdmin/internal/controller.(*cAuthRule).Update
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | DELETE | /org                | GaAdmin/internal/controller.(*cOrg).Delete
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /org                | GaAdmin/internal/controller.(*cOrg).Get
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /org                | GaAdmin/internal/controller.(*cOrg).Create
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /org                | GaAdmin/internal/controller.(*cOrg).Update
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | DELETE | /org/member         | GaAdmin/internal/controller.(*cOrgMember).Delete
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /org/member         | GaAdmin/internal/controller.(*cOrgMember).Get
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /org/member         | GaAdmin/internal/controller.(*cOrgMember).Create
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /org/member         | GaAdmin/internal/controller.(*cOrgMember).Update
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /org/members        | GaAdmin/internal/controller.(*cOrgMember).List
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /orgs               | GaAdmin/internal/controller.(*cOrg).List
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | DELETE | /user               | GaAdmin/internal/controller.(*cUser).Delete
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /user               | GaAdmin/internal/controller.(*cUser).Get
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user               | GaAdmin/internal/controller.(*cUser).Create
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /user               | GaAdmin/internal/controller.(*cUser).Update
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | DELETE | /user/group         | GaAdmin/internal/controller.(*cUserGroup).Delete
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /user/group         | GaAdmin/internal/controller.(*cUserGroup).Get
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user/group         | GaAdmin/internal/controller.(*cUserGroup).Create
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /user/group         | GaAdmin/internal/controller.(*cUserGroup).Update
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user/group-access  | GaAdmin/internal/controller.(*cUserGroupAccess).Setup
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /user/group-tree    | GaAdmin/internal/controller.(*cUserGroup).Tree
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user/refresh-token | GaAdmin/internal/controller.(*cUserSign).Refresh
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user/sign-mobile   | GaAdmin/internal/controller.(*cUserSign).SignMobile
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | PUT    | /user/sign-out      | GaAdmin/internal/controller.(*cUserSign).SignOut
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user/sign-passport | GaAdmin/internal/controller.(*cUserSign).SignPassport
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | POST   | /user/sign-up       | GaAdmin/internal/controller.(*cUserSign).SignUp
----------|--------|---------------------|-----------------------------------------------------------------
  :8000   | GET    | /users              | GaAdmin/internal/controller.(*cUser).List
----------|--------|---------------------|----------------------------------------------------------------- 

```

# 特别鸣谢
- GoFrame https://github.com/gogf/gf
- Ant design https://github.com/ant-design/ant-design
- Casbin https://github.com/casbin/casbin
- GolangJwt https://github.com/golang-jwt/jwt

# 版权信息
GaAdmin遵循 [MIT license](LICENSE) 开源协议发布，并提供免费使用。



