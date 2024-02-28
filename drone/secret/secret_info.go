package secret

import (
	"errors"
	"html/template"
	"os"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/drone/funcmap"
	"github.com/urfave/cli"
)

var secretInfoCmd = cli.Command{
	Name:      "info",
	Usage:     "display secret info",
	ArgsUsage: "[repo/name]",
	Action:    secretInfo,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "repository",
			Usage: "repository name (e.g. octocat/hello-world)",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "secret name",
		},
		cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: tmplSecretList,
		},
	},
}

func secretInfo(c *cli.Context) error {
	var (
		secretName = c.String("name")
		repoName   = c.String("repository")
		format     = c.String("format") + "\n"
	)
	if secretName == "" {
		return errors.New("Missing secret name")
	}
	if repoName == "" {
		repoName = c.Args().First()
	}
	owner, name, err := internal.ParseRepo(repoName)
	if err != nil {
		return err
	}
	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}
	secret, err := client.Secret(owner, name, secretName)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(format)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, secret)
}
