{{ template "header.tpl" . }}

<h1>{{.Text}}</h1>
{{range $key, $val := .s}}
{{$key}}
{{$val}}
{{end}}
{{ template "footer.tpl" . }}