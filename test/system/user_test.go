package system

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhanghup/go-app/test"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestUserTemplate(t *testing.T) {
	test.MainTest(t, "user", map[string]interface{}{
		"type":     "a",
		"account":  tools.ObjectString(),
		"password": "Aa123456",
		"name":     "张超",
		"avatar":   "www.baidu.com",
		"i_card":   "330411199206123611",
		"birth":    1567252216,
		"sex":      1,
		"mobile":   "15115151515",
		"admin":    1,
		"weight":   0,
		"status":   1,
	})
}

func TestUserCreate(t *testing.T) {
	id := tools.ObjectString()
	result := map[string]interface{}{}
	_, err := test.Query(t, `
		mutation UserCreate($input:NewUser!){
			user_create(input:$input){
				id
				type
				account
				password
				name
				avatar
				i_card
				sex
				mobile
				admin
				weight
				status
				created
				updated
			}
		}
	`, map[string]interface{}{
		"input": map[string]interface{}{
			"type":     "a",
			"account":  *id,
			"password": "Aa123456",
			"name":     "张超",
			"avatar":   "www.baidu.com",
			"i_card":   "330411199206123611",
			"birth":    1567252216,
			"sex":      1,
			"mobile":   "15115151515",
			"admin":    1,
			"weight":   0,
			"status":   1,
		},
	}, &result)
	assert.NoError(t, err)
	newId := ((result["user_create"]).(map[string]interface{})["id"]).(string)

	result = map[string]interface{}{}
	_, err = test.Query(t, `
		mutation UserUpdate($id:String!,$input:UpdUser!){
			user_update(id:$id,input:$input)
		}
	`, map[string]interface{}{
		"id": newId,
		"input": map[string]interface{}{
			"type":     "a",
			"account":  id,
			"password": "Aa123456",
			"name":     "张超222",
			"avatar":   "www.baidu.com",
			"i_card":   "330411199206123611",
			"birth":    1567252216,
			"sex":      1,
			"mobile":   "15115151515",
			"admin":    1,
			"weight":   0,
			"status":   1,
		},
	}, &result)
	assert.NoError(t, err)

	_, err = test.Query(t, `
		mutation UserUpdate($id:[String!]){
			user_removes(ids:$id)
		}
	`, map[string]interface{}{
		"id": []string{newId},
	}, &map[string]interface{}{})
	assert.NoError(t, err)
}
