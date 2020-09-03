package cmd

// AppHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
// https://github.com/urfave/cli/blob/master/help.go
var appHelpTemplate = `NAME:
   {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

USAGE:
{{- if .UsageText }}
	 {{.UsageText}}
{{- else }}
	{{.HelpName}}
	{{if .VisibleFlags}}
		[global options]
	{{end}}
	{{if .Commands}}
		command [command options]
	{{end}}
	{{if .ArgsUsage}}
		{{.ArgsUsage}}
	{{else}}
		[arguments...]
	{{end}}
{{- end}}


{{- if .VisibleFlags}}

OPTIONS:
	{{- range $index, $option := .VisibleFlags }}
		{{$option}}
	{{- end -}}
{{- end }}

EXAMPLES:
	- start server at http://0.0.0.0:9000 serving "." current directory
			statiks -port 9000

	- start server at http://0.0.0.0:9080 serving "/home" with CORS
			statiks --cors /home

	- start server at http://192.168.1.100:9080 serving "/tmp" with gzip compression
			statiks --host 192.168.1.100 --compression /tmp

	- start server at https://0.0.0.0:9080 serving "." with HTTPS
			statiks --ssl --cert cert.pem --key key.pem

	- start server at http://0.0.0.0:9080 serving "/tmp" with delay response 100ms
			statiks -add-delay 100 /tmp

{{- if .Version }}

VERSION:
   {{ .Version }}
{{- end}}

{{- if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}
{{end}}
`
