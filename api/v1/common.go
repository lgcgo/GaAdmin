package v1

import "github.com/lgcgo/tree"

type Page struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Order string `json:"order"`
}

type Pager struct {
	Page  int `json:"currentPage"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

type TreeResData struct {
	TreeData *tree.TreeData `json:"treeData"`
	Total    int            `json:"total"`
}
