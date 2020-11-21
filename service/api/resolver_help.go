package api

import "github.com/zhanghup/go-app/beans"

func (r *queryResolver) DeptTreeHelp(items []beans.Dept, pid string, flag ...bool) interface{} {
	type DeptTreeItem struct {
		Id       *string     `json:"id"`
		Name     *string     `json:"name"`
		Code     *string     `json:"code"`
		Children interface{} `json:"children"`
	}

	results := make([]DeptTreeItem, 0)
	for _, o := range items {
		item := DeptTreeItem{
			Id:   o.Id,
			Name: o.Name,
			Code: o.Code,
		}

		if len(flag) > 0 && flag[0] {
			if o.Pid == nil {
				item.Children = r.DeptTreeHelp(items, *o.Id)
				results = append(results, item)
			}
		} else {
			if *o.Pid == pid {
				item.Children = r.DeptTreeHelp(items, *o.Id)
				results = append(results, item)
			}
		}
	}
	return results
}
