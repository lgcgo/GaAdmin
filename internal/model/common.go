package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lgcgo/tree"
)

// 数据分页
type Page struct {
	Page      int
	Size      int
	Order     string
	Condition g.Map
}
type Pager struct {
	Page  int
	Size  int
	Total int
}

/**
* 树状数据
**/
type TreeDataOutput struct {
	TreeData *tree.TreeData
	Total    uint
	Keys     []uint
}
