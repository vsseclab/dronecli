package server

import (
	"fmt"
	"os"
	"text/template"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/drone/funcmap"
	"github.com/urfave/cli"
)

var serverInfoCmd = cli.Command{
	Name:      "info",
	Usage:     "show server details",
	ArgsUsage: "<servername>",
	Action:    serverInfo,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: tmplServerInfo,
		},
	},
}

func serverInfo(c *cli.Context) error {
	client, err := internal.NewAutoscaleClient(c)
	if err != nil {
		return err
	}

	name := c.Args().First()
	if len(name) == 0 {
		return fmt.Errorf("Missing or invalid server name")
	}

	server, err := client.Server(name)
	if err != nil {
		return err
	}

	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.String("format") + "\n")
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, server)
}

// template for server information
var tmplServerInfo = `Name: {{ .Name }}
Address: {{ .Address }}
Region:  {{ .Region }}
Size:    {{.Size}}
State:   {{ .State }}
{{ if .Error -}}
Error:   {{ .Error }}
{{ end -}}
`
