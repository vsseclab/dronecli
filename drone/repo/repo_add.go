package repo

import (
	"fmt"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var repoAddCmd = cli.Command{
	Name:      "enable",
	Usage:     "enable a repository",
	ArgsUsage: "<repo/name>",
	Action:    repoAdd,
}

func repoAdd(c *cli.Context) error {
	repo := c.Args().First()
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	if _, err := client.RepoEnable(owner, name); err != nil {
		return err
	}
	fmt.Printf("Successfully activated repository %s/%s\n", owner, name)
	return nil
}
