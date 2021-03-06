## 介绍
GaAdmin是基于GoFrame+Antd的中后台开发框架。系统集成了中后台管理系统通用的基础功能模块，并提供实用的代码生成工具，能够大幅降低编写通用的代码，使开发者专注于应用的业务开发。

## 主要特性
- 统一通用的用户与鉴权，支持多应用下的单点登录
- 基于 Casbin + Jwt 高效的 RBAC 权限认证模块
- 命令一键生成 CURD+ 接口以及服务层业务逻辑
- 接口遵循 OpenAPIv3 协议，自动构建Swagger文档

## 内置功能
- **用户模块**：用户分组，用户管理，支持用户多分组设置。
- **权限菜单**：设置菜单与权限规则，按用户组进行权限划分。
- **组织机构**：部门小组，组织成员，支持树状结构展示。
- **系统设置**：配置系统基础信息、常用参数（开发中）
- **消息队列**：支持Mysql、Readis、Kafka等缓存方案（开发中）
- **通知公告**：消息队列群发用户公告（开发中）
- **系统日志**：系统操作、系统异常信息日志记录（开发中）
- 更多功能正在积极开发中...

## 接口清单

```shell
 ADDRESS | METHOD |             ROUTE             |                           HANDLER                        
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | DELETE | /auth/menu                    | GaAdmin/internal/controller.(*cAuthMenu).Delete          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /auth/menu                    | GaAdmin/internal/controller.(*cAuthMenu).Get             
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /auth/menu                    | GaAdmin/internal/controller.(*cAuthMenu).Create          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /auth/menu                    | GaAdmin/internal/controller.(*cAuthMenu).Update          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /auth/role-access             | GaAdmin/internal/controller.(*cAuthRoleAccess).Setup     
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /auth/menu-tree               | GaAdmin/internal/controller.(*cAuthMenu).Tree            
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | DELETE | /auth/role                    | GaAdmin/internal/controller.(*cAuthRole).Delete          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /auth/role                    | GaAdmin/internal/controller.(*cAuthRole).Get             
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /auth/role                    | GaAdmin/internal/controller.(*cAuthRole).Create          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /auth/role                    | GaAdmin/internal/controller.(*cAuthRole).Update          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /auth/role-tree               | GaAdmin/internal/controller.(*cAuthRole).Tree            
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | DELETE | /auth/rule                    | GaAdmin/internal/controller.(*cAuthRule).Delete          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /auth/rule                    | GaAdmin/internal/controller.(*cAuthRule).Get             
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /auth/rule                    | GaAdmin/internal/controller.(*cAuthRule).Create          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /auth/rule                    | GaAdmin/internal/controller.(*cAuthRule).Update          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | DELETE | /org/department               | GaAdmin/internal/controller.(*cOrgDepartment).Delete     
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /org/department               | GaAdmin/internal/controller.(*cOrgDepartment).Get        
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /org/department               | GaAdmin/internal/controller.(*cOrgDepartment).Create     
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /org/department               | GaAdmin/internal/controller.(*cOrgDepartment).Update     
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /org/department-tree          | GaAdmin/internal/controller.(*cOrgDepartment).Tree       
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | DELETE | /org/member                   | GaAdmin/internal/controller.(*cOrgMember).Delete         
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /org/member                   | GaAdmin/internal/controller.(*cOrgMember).Get            
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /org/member                   | GaAdmin/internal/controller.(*cOrgMember).Create         
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /org/member                   | GaAdmin/internal/controller.(*cOrgMember).Update         
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /org/members                  | GaAdmin/internal/controller.(*cOrgMember).List           
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | DELETE | /user                         | GaAdmin/internal/controller.(*cUser).Delete              
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /user                         | GaAdmin/internal/controller.(*cUser).Get                 
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user                         | GaAdmin/internal/controller.(*cUser).Create              
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /user                         | GaAdmin/internal/controller.(*cUser).Update              
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/password/reset-email    | GaAdmin/internal/controller.(*cUserReset).ResetEmail     
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/password/reset-mobile   | GaAdmin/internal/controller.(*cUserReset).ResetMobile    
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/password/reset-question | GaAdmin/internal/controller.(*cUserReset).ResetQuestion  
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/refresh-token           | GaAdmin/internal/controller.(*cUserSign).Refresh         
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/roles                   | GaAdmin/internal/controller.(*cUserAccess).Setup         
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/sign-mobile             | GaAdmin/internal/controller.(*cUserSign).SignMobile      
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | PUT    | /user/sign-out                | GaAdmin/internal/controller.(*cUserSign).SignOut         
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/sign-passport           | GaAdmin/internal/controller.(*cUserSign).SignPassport    
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | POST   | /user/sign-up                 | GaAdmin/internal/controller.(*cUserSign).SignUp          
---------|--------|-------------------------------|----------------------------------------------------------
  :8000  | GET    | /users                        | GaAdmin/internal/controller.(*cUser).List                
---------|--------|-------------------------------|----------------------------------------------------------

```

## 前端项目
- 待完善...

## 安装使用

想要提前体验或者学习参考的同学，可以按以下步骤启动项目进行测试。

**Step1**: 下载源码
```
git clone https://github.com/lgcgo/GaAdmin
```

**Step2**: 导入数据
- 创建Mysql数据库gaadmin
- 导入manifest/data/data.sql

**Step3**: 启动项目

命令进入根目录
```Shell
cd GaAdmin
```

命令启动
```Golang
go run .
```

**Step4**: 开始测试

- 接口文档：http://127.0.0.1:8000/swagger/

- 测试工具：推荐使用Apifox，URL导入http://127.0.0.1:8000/api.json 开始测试

- 测试用户：账号admin 密码123456

## 环境要求 
- 待完善...

## 技术文档
- http://doc.gaadmin.com(筹备中)

## 特别鸣谢
- GoFrame https://github.com/gogf/gf
- Ant design https://github.com/ant-design/ant-design
- Casbin https://github.com/casbin/casbin
- GolangJwt https://github.com/golang-jwt/jwt

## 版权信息
GaAdmin遵循 [MIT license](LICENSE) 开源协议发布，并提供免费使用。



