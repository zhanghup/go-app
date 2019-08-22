package test

import (
	"fmt"
	"testing"
)

func TestQueryHello(t *testing.T) {
	r,err := query(t,`
		query S{
			hello	
		}
	`,nil)
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Body.String())

}
