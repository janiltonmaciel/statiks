package core

import (
	"fmt"

	"github.com/urfave/cli"
)

// https://github.com/urfave/cli/blob/master/help.go
// AppHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var AppHelpTemplate = `NAME:
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

EXAMPLES
	- start server at http://localhost:9000 serving "." with allowed transfer of hidden files
			statiks -p 9000 --hidden

	- start server at http://localhost:9080 serving "~/Projects" with allowed methods "GET, POST"
			statiks -m "GET, POST" ~/Projects

	- start server at http://192.168.1.100:9080 serving "~/Data" with enabled gzip compression
			statiks -t 192.168.1.100 --compress ~/Data


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

func VersionPrinter(commit, date string) func(c *cli.Context) {
	return func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version: %s\n", c.App.Version)
		fmt.Fprintf(c.App.Writer, "author: %s\n", c.App.Author)
		fmt.Fprintf(c.App.Writer, "commit: %s\n", commit)
		fmt.Fprintf(c.App.Writer, "date: %s\n", date)
	}
}
