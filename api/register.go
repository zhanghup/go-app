package api

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

//var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
//	Query:    rootQuery,
//	Mutation: rootMutation, // 需要通过GraphQL更新数据，可以定义Mutation
//})

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// 只需要通过Gin简单封装即可
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// 定义跟查询节点
//var rootQuery = graphql.NewObject(graphql.ObjectConfig{
//	Name:        "RootQuery",
//	Description: "数据查询",
//	Fields: graphql.Fields{
//		"login": &Login,
//	},
//})
//
//// 定义跟查询节点
//var rootMutation = graphql.NewObject(graphql.ObjectConfig{
//	Name:        "RootMutation",
//	Description: "数据操作",
//	Fields: graphql.Fields{
//		"login": &Login,
//	},
//})
//
//func RegistQuery(fieldName string, fieldConfig *graphql.Field) {
//	rootQuery.AddFieldConfig(fieldName, fieldConfig)
//}
//func RegistMustation(fieldName string, fieldConfig *graphql.Field) {
//	rootMutation.AddFieldConfig(fieldName, fieldConfig)
//}
