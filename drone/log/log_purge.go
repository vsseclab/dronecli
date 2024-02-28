package log

import (
	"fmt"
	"strconv"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var logPurgeCmd = cli.Command{
	Name:      "purge",
	Usage:     "purge a log",
	ArgsUsage: "<repo/name> <build> <stage> <step>",
	Action:    logPurge,
}

func logPurge(c *cli.Context) (err error) {
	repo := c.Args().First()
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}
	number, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return err
	}
	stage, err := strconv.Atoi(c.Args().Get(2))
	if err != nil {
		return err
	}
	step, err := strconv.Atoi(c.Args().Get(3))
	if err != nil {
		return err
	}

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	err = client.LogsPurge(owner, name, number, stage, step)
	if err != nil {
		return err
	}

	fmt.Printf("Purging logs for build %s/%s#%d\n", owner, name, number)
	return nil
}
