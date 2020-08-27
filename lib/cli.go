package lib

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// AppHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
// https://github.com/urfave/cli/blob/master/help.go
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

EXAMPLES:
	- start server at http://0.0.0.0:9000 serving "." current directory
			statiks -port 9000

	- start server at http://0.0.0.0:9080 serving "/home" with CORS
			statiks --cors /home

	- start server at http://192.168.1.100:9080 serving "/tmp" with gzip compression
			statiks --host 192.168.1.100 --compression /tmp

	- start server at https://0.0.0.0:9080 serving "." with HTTPS
			statiks --ssl --cert cert.pem --key key.pem

	- start server at http://0.0.0.0:9000 serving "/tmp" with delay response 100ms
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

func VersionPrinter(commit, date string) func(c *cli.Context) {
	return func(c *cli.Context) {
		printLogo()
		fmt.Fprintf(c.App.Writer, "version: %s\n", c.App.Version)
		fmt.Fprintf(c.App.Writer, "commit: %s\n", commit)
		fmt.Fprintf(c.App.Writer, "date: %s\n", date)
		fmt.Fprintf(c.App.Writer, "author: %s\n", c.App.Authors[0].Name)
	}
}

func printLogo() {
	fmt.Println(`
     _______.___________.    ___   .___________. __   __  ___      _______.
    /       |           |   /   \  |           ||  | |  |/  /     /       |
   |   (----'---|  |----'  /  ^  \ '---|  |----'|  | |  '  /     |   (----'
    \   \       |  |      /  /_\  \    |  |     |  | |    <       \   \
.----)   |      |  |     /  _____  \   |  |     |  | |  .  \  .----)   |
|_______/       |__|    /__/     \__\  |__|     |__| |__|\__\ |_______/`)
}
