package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// 定义Schema用于http handler处理
type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func main() {
	s := `
                schema {
                        query: Query
                }
                type Query {
                        hello: String!
                }
        `

	schema := graphql.MustParseSchema(s, &query{})

	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		var page = []byte(`
		<!DOCTYPE html>
		<html>
			<head>
				<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
				<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
			</head>
			<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
				<div id="graphiql" style="height: 100vh;">Loading...</div>
				<script>
					function graphQLFetcher(graphQLParams) {
						return fetch("/query", {
							method: "post",
							body: JSON.stringify(graphQLParams),
							credentials: "include",
						}).then(function (response) {
							return response.text();
						}).then(function (responseBody) {
							try {
								return JSON.parse(responseBody);
							} catch (error) {
								return responseBody;
							}
						});
					}
					ReactDOM.render(
						React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
						document.getElementById("graphiql")
					);
				</script>
			</body>
		</html>
		`)
		c.Writer.Write(page)
	})
	g.POST("/query", func(c *gin.Context) {
		hunder := &relay.Handler{Schema: schema}
		hunder.ServeHTTP(c.Writer, c.Request)
	})
	g.Run(":8899")

}
