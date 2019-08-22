package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zhanghup/go-app/api/server/engine"
	"github.com/zhanghup/go-tools"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
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
	req, err := http.NewRequest("POST", prefix, strings.NewReader(tools.JSONString(q)))
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
	err = json.Unmarshal([]byte(w.Body.String()), &obj)
	if err != nil {
		assert.Error(t, err, "json解析异常")
		return w, err
	}

	for _, v := range obj {
		err = json.Unmarshal([]byte(tools.JSONString(v)), result[0])
		if err != nil {
			assert.Error(t, err, "json解析异常")
		}
		break
	}

	return w, err
}

func query(t assert.TestingT, query string, variables interface{}, result ...interface{}) (*httptest.ResponseRecorder, error) {
	return action(t, "/query", query, variables, result...)
}

var e *gin.Engine

func init() {
	e = engine.Router()
	//login(e)
}
