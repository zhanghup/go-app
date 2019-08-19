package api

//import "github.com/graphql-go/graphql"

// 处理查询请求
//var Login = graphql.Field{
//	Description: "web用户登录",
//	Type:        graphql.String,
//	Args: graphql.FieldConfigArgument{
//		"account": &graphql.ArgumentConfig{
//			Type: graphql.String,
//		},
//		"password": &graphql.ArgumentConfig{
//			Type: graphql.String,
//		},
//	},
//	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
//		id, _ := p.Args["account"].(string)
//		name, _ := p.Args["password"].(string)
//
//		return id + "," + name, nil
//	},
//}
