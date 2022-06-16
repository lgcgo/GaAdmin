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
  ADDRESS | METHOD |       ROUTE        |                             HANDLER                             |     MIDDLEWARE       
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | ALL    | /*                 | github.com/gogf/gf/v2/net/ghttp.internalMiddlewareServerTracing | GLOBAL MIDDLEWARE    
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | ALL    | /api.json          | github.com/gogf/gf/v2/net/ghttp.(*Server).openapiSpec           |
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | ALL    | /swagger/*         | github.com/gogf/gf/v2/net/ghttp.(*Server).swaggerUI             | HOOK_BEFORE_SERVE
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | DELETE | /user              | GaAdmin/internal/controller.(*cUser).Delete                     | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | GET    | /user              | GaAdmin/internal/controller.(*cUser).Get                        | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | POST   | /user              | GaAdmin/internal/controller.(*cUser).Create                     | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | PUT    | /user              | GaAdmin/internal/controller.(*cUser).Update                     | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | POST   | /user/group-access | GaAdmin/internal/controller.(*cUserGroupAccess).Setup           | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | GET    | /user/group-tree   | GaAdmin/internal/controller.(*cUserGroup).Tree                  | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
  :8000   | GET    | /users             | GaAdmin/internal/controller.(*cUser).List                       | middleware.CORS
          |        |                    |                                                                 | middleware.Context
          |        |                    |                                                                 | middleware.Response
----------|--------|--------------------|-----------------------------------------------------------------|----------------------
```

# 特别鸣谢
- GoFrame https://github.com/gogf/gf
- Ant design https://github.com/ant-design/ant-design
- Casbin https://github.com/casbin/casbin
- GolangJwt https://github.com/golang-jwt/jwt

# 版权信息
GaAdmin遵循 [MIT license](LICENSE) 开源协议发布，并提供免费使用。



