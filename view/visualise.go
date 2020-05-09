package view

import (
	"bytes"
	"github.com/gofiber/fiber"
	"html/template"
)

var Visualize = `<!DOCTYPE html>
<html>
  <head>
    <script src="//cdn.jsdelivr.net/npm/react@16/umd/react.production.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/fetch/3.0.0/fetch.min.js"></script>
    <script src="//cdn.jsdelivr.net/npm/react-dom@16/umd/react-dom.production.min.js"></script>

    <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-voyager/dist/voyager.css" />
    <script src="//cdn.jsdelivr.net/npm/graphql-voyager/dist/voyager.min.js"></script>
  </head>
  <body>
    <div id="voyager">Loading...</div>
    <script>
      function introspectionProvider(query) {
        // ... do a call to server using introspectionQuery provided
        // or just return pre-fetched introspection
		 return window.fetch(window.location.origin + '/query', {
            method: 'post',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({query}),
            }).then(response => response.json());
      }

      // Render <Voyager />
      GraphQLVoyager.init(document.getElementById('voyager'), {
        introspection: introspectionProvider
      })
    </script>
	<style>
	.powered-by {
    display: none !important;
	}
div#voyager {
    height: 100%;
    position: absolute;
}
	</style>
  </body>
</html>
`

var v = template.Must(template.New("html").Parse(Visualize))

func Visualise(title string, endpoint string) []byte {
	body := new(bytes.Buffer)
	_ = v.Execute(body, map[string]string{
		"title":    title,
		"endpoint": endpoint,
	})
	return body.Bytes()
}

func MountVisualDependencyGraph(c *fiber.Ctx) {
	c.Set("content-type", "text/html")
	c.SendBytes(Visualise("Something", "/query"))
}
