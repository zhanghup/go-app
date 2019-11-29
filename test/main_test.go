package test

import (
	"fmt"
	"github.com/zhanghup/go-app/libs/wx/qiye"
	"os"
	"testing"
	"text/template"
)

func TestName(t *testing.T) {
	tt := template.New("member")
	tt, err := tt.Parse("Member named {{ .Name}} with description: {{ .Description}}")
	if err != nil {
		panic(err)
	}
	err = tt.Execute(os.Stdout, map[string]interface{}{
		"Name":        "123",
		"Description": "Description222",
	})
	if err != nil {
		panic(err)
	}
}



func TestQueryHello(t *testing.T) {
	r, err := Query(t, `
		query S{
			hello	
		}
	`, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Body.String())

}
