package gs

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Gin(g context.Context) *gin.Context {
	gg := g.Value(GIN_CONTEXT)
	return gg.(*gin.Context)
}

func GinPlayground(g gin.IRouter, url, query string) {
	g.GET(url, func(c *gin.Context) {
		var page = []byte(fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
			<head>
				<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
				<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
			</head>
			<body style="width: 100%%; height: 100%%; margin: 0; overflow: hidden;">
				<div id="graphiql" style="height: 100vh;">Loading...</div>
				<script>
					function graphQLFetcher(graphQLParams) {
						return fetch("%s", {
							method: "post",
							headers:{
								"Content-Type": "application/json"
							},
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
		`, query))
		_, _ = c.Writer.Write(page)
	})
}
