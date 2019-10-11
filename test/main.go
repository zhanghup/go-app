package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/zhanghup/go-app/api/server/engine"
	"github.com/zhanghup/go-tools"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"text/template"
)

var cookies []*http.Cookie

func login(g *gin.Engine) {
	//r, err := Iot(nil, `
	//	mutation S {
	//	  login(account: "admin", password: "junji")
	//	}
	//`, nil)
	//
	//if err != nil {
	//	panic(err)
	//}
	//cookies = r.Result().Cookies()
}

func GetOperationName(str string) string {
	r, _ := regexp.Compile(`(mutation|query)\s(.*?)\s{`)
	rr := r.FindStringSubmatch(str)
	if len(rr) >= 3 {
		return rr[2]
	}
	return ""
}

func action(t assert.TestingT, prefix, query string, variables interface{}, result ...interface{}) (*httptest.ResponseRecorder, error) {
	q := map[string]interface{}{
		"operationName": GetOperationName(query),
		"query":         query,
		"variables":     variables,
	}
	req, err := http.NewRequest("POST", prefix, strings.NewReader(tools.Str().JSONString(q)))
	if err != nil {
		panic(err)
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	if result == nil || len(result) == 0 {
		return w, nil
	}
	obj := map[string]interface{}{}
	fmt.Println(w.Body.String())
	err = json.Unmarshal([]byte(w.Body.String()), &obj)
	if err != nil {
		assert.Error(t, err, "json解析异常")
		return w, err
	}
	if o, ok := obj["data"]; !ok || o == nil {
		assert.Error(t, err, "graphql数据异常")
		t.Errorf("graphql数据异常:", w.Body.String())
		return w, err
	}
	if _, ok := obj["errors"]; ok {
		assert.Error(t, err, "graphql数据异常")
		t.Errorf("graphql数据异常:", w.Body.String())
		return w, err
	}

	for _, v := range obj {
		err = json.Unmarshal([]byte(tools.Str().JSONString(v)), result[0])
		if err != nil {
			assert.Error(t, err, "json解析异常")
		}

		break
	}

	return w, err
}

func Query(t assert.TestingT, query string, variables interface{}, result ...interface{}) (*httptest.ResponseRecorder, error) {
	return action(t, "/base", query, variables, result...)
}

func Tpl() template.FuncMap {
	return template.FuncMap{
		"title": func(str string) string {
			ss := strings.Split(str, "_")
			for i := range ss {
				ss[i] = strings.Title(ss[i])
			}

			return strings.Join(ss, "")
		},
	}
}

type Mt struct {
	create bool
	update bool
	remove bool
	query  bool
	get    bool
}
type MtType string

const (
	MtCreate = MtType("1")
	MtUpdate = MtType("2")
	MtRemove = MtType("3")
	MtQuery  = MtType("4")
	MtGet    = MtType("5")
)

func NewMt(ty ...MtType) Mt {
	r := Mt{
		create: true,
		update: true,
		remove: true,
		query:  true,
		get:    true,
	}
	for _, t := range ty {
		switch t {
		case MtType("1"):
			r.create = false
		case MtType("2"):
			r.update = false
		case MtType("3"):
			r.remove = false
		case MtType("4"):
			r.query = false
		case MtType("5"):
			r.get = false
		}
	}
	return r
}

func (this Mt) MainTest(t *testing.T, obj string, params ...map[string]interface{}) {
	shows := ""
	var createParam, updateParam map[string]interface{}
	if len(params) > 0 {
		createParam = params[0]
		updateParam = params[0]
		if len(params) > 1 {
			updateParam = params[1]
		}
		for k := range createParam {
			shows += fmt.Sprintf("\t\t\t\t %s\n", k)
		}
		if _, ok := createParam["id"]; !ok {
			shows = " id \n " + shows
		}
		if _, ok := createParam["weight"]; !ok {
			shows += "\t\t\t\t weight \n "
		}
		if _, ok := createParam["status"]; !ok {
			shows += "\t\t\t\t status \n "
		}
		if _, ok := createParam["created"]; !ok {
			shows += "\t\t\t\t created \n "
		}
		if _, ok := createParam["updated"]; !ok {
			shows += "\t\t\t\t updated"
		}
	}

	// graphql 新增
	create, err := tools.Str().Template(`
		mutation {{title .object -}}Create($input:New{{- title .object -}}!){
			{{ .object -}}_create(input:$input){
				{{.shows}}
			}
		}
	`, map[string]interface{}{
		"object": obj,
		"shows":  shows,
	}, Tpl())
	assert.NoError(t, err, err)

	query, err := tools.Str().Template(`
		query {{ title .object -}}Query($query:Q{{- title .object -}}!){
			{{ .object -}}s(query:$query){
				total
				data{
				{{.shows}}
				}
			}
		}
	`, map[string]interface{}{
		"object": obj,
		"shows":  shows,
	}, Tpl())
	assert.NoError(t, err, err)

	get, err := tools.Str().Template(`
		query {{ title .object -}}Get($id:String!){
			{{ .object -}}(id:$id){
				{{.shows}}
			}
		}
	`, map[string]interface{}{
		"object": obj,
		"shows":  shows,
	}, Tpl())
	assert.NoError(t, err, err)

	// graphql 修改
	update, err := tools.Str().Template(`
		mutation {{ title .object -}}Update($id:String!,$input:Upd{{- title .object -}}!){
			{{.object -}}_update(id:$id,input:$input)
		}
	`, map[string]interface{}{"object": obj}, Tpl())
	assert.NoError(t, err, err)

	// graphql 删除
	remove, err := tools.Str().Template(`
		mutation {{title .object -}}Remove($id:[String!]){
			{{.object -}}_removes(ids:$id)
		}
	`, map[string]interface{}{
		"object": obj,
	}, Tpl())
	assert.NoError(t, err, err)

	// 新增
	result := map[string]interface{}{}
	if this.create {
		_, err = Query(t, create, map[string]interface{}{"input": createParam}, &result)
		assert.NoError(t, err)
	}
	newId := ((result[obj+"_create"]).(map[string]interface{})["id"]).(string)

	result = map[string]interface{}{}

	// 批量查询
	if this.query {
		result = map[string]interface{}{}
		_, err = Query(t, query, map[string]interface{}{
			"query": map[string]interface{}{"count": true, "size": 2},
		}, &result)
		assert.NoError(t, err)
	}

	// 单个查询
	if this.get {
		result = map[string]interface{}{}
		_, err = Query(t, get, map[string]interface{}{
			"id": newId,
		}, &result)
		assert.NoError(t, err)
	}

	// 更新
	if this.update {
		result = map[string]interface{}{}
		_, err = Query(t, update, map[string]interface{}{
			"id":    newId,
			"input": updateParam,
		}, &result)
		assert.NoError(t, err)
	}

	// 删除
	if this.remove {
		result = map[string]interface{}{}
		_, err = Query(t, remove, map[string]interface{}{
			"id": []string{newId},
		}, &map[string]interface{}{})
		assert.NoError(t, err)
	}

	// 单个查询
	if this.get {
		result = map[string]interface{}{}
		_, err = Query(t, get, map[string]interface{}{
			"id": newId,
		}, &result)
		assert.NoError(t, err)
		o, ok := (result[obj]).(map[string]interface{})["id"]
		if ok && o != nil {
			t.Error(errors.New("删除异常，没有删除掉"), tools.Str().JSONString(result), remove)
		}
	}
}

var e *gin.Engine

func init() {
	e = engine.Router()
	//login(e)
}
