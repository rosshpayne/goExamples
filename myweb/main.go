
package main

import (
  "html/template"
  "net/http"
)

var tmpl = `<html>
<head>
    <title>Hello World!</title>
</head>
<body>
    {{ . }}
</body>
</html>
`
func handler(w http.ResponseWriter, r *http.Request) {
  t := template.New("main") //name of the template is main
  t, _ = t.Parse(tmpl) // parsing of template string
  t.Execute(w, "Hello World again!")
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:9003",
  }
  http.HandleFunc("/view", handler)
  server.ListenAndServe()
}
