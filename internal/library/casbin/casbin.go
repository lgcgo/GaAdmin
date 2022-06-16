package casbin

import (
	"bufio"
	"errors"
	"os"
	"strings"

	pkg "github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type Casbin struct {
	E              *pkg.Enforcer
	ModelFilePath  string
	PolicyFilePath string
}

// Plicy 代表访问节点
type Policy struct {
	Subject string `json:"sub" v:"required"` // 代表用户组名
	Object  string `json:"obj" v:"required"` // 代表请求路径
	Action  string `json:"act" v:"required"` // 代表请求方法
}

// Role 代表用户组关系
type Role struct {
	ParentSubject string
	Subject       string
}

var insCasbin *Casbin

func NewCasbin() *Casbin {
	var (
		ctx            = gctx.New()
		modelFilePath  = g.Cfg().MustGet(ctx, "casbin.modelFilePath").String()
		policyFilePath = g.Cfg().MustGet(ctx, "casbin.policyFilePath").String()
		enforce        *pkg.Enforcer
		err            error
	)

	enforce, err = pkg.NewEnforcer(modelFilePath, policyFilePath)
	if err != nil {
		panic(err.Error())
	}
	enforce.EnableAutoSave(true)
	insCasbin = &Casbin{
		enforce,
		modelFilePath,
		policyFilePath,
	}
	return insCasbin
}

func (p *Policy) Format() string {
	var (
		strArr []string
	)
	strArr = append(strArr, "p")
	strArr = append(strArr, p.Subject)
	strArr = append(strArr, p.Object)
	strArr = append(strArr, p.Action)
	return strings.Join(strArr, ", ")
}

func (r *Role) Format() string {
	var (
		strArr []string
	)
	strArr = append(strArr, "g")
	strArr = append(strArr, r.ParentSubject)
	strArr = append(strArr, r.Subject)
	return strings.Join(strArr, ", ")
}

// 检测Policy
func (c *Casbin) Verify(p *Policy) (bool, error) {
	var (
		err error
		ok  bool
	)
	if ok, err = c.E.Enforce(p.Subject, p.Object, p.Action); err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("deny the request")
	}
	return true, nil
}

// 更新Policy.csv
func (c *Casbin) SavePolicyCsv(policys []*Policy, roles []*Role) error {
	var (
		file   *os.File
		writer *bufio.Writer
		err    error
	)
	// 获取文件句柄
	file, err = os.OpenFile(c.PolicyFilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	// 写入权限节点
	writer = bufio.NewWriter(file)
	for _, policy := range policys {
		writer.WriteString(policy.Format())
		writer.WriteString("\n")
	}
	// 写入角色关系
	for _, role := range roles {
		writer.WriteString(role.Format())
		writer.WriteString("\n")
	}
	writer.Flush()
	return nil
}
