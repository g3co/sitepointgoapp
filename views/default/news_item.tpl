{{ template "header.tpl" . }}

<h1>{{.Title}}</h1>

{{.Text | str2html}}

{{ template "footer.tpl" . }}