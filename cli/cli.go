package cli

import (
	"os"

	"github.com/jvikstedt/alarmii/cli/action"
	"github.com/jvikstedt/alarmii/cli/service"
	"github.com/urfave/cli"
)

var jobAction = action.NewJob(os.Stdout, service.Job{})
var actionMap = map[string]action.Execute{
	"job_add":  jobAction.Add,
	"job_list": jobAction.List,
}

// Run executes cli app
func Run() {
	app := cli.NewApp()
	app.Name = "Alarmii"
	app.Usage = ""

	setupActions(app)

	app.Run(os.Args)
}

func setupActions(app *cli.App) {
	app.Commands = []cli.Command{
		{
			Name: "job",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Action: getAction("job_add"),
				},
				{
					Name:   "list",
					Action: getAction("job_list"),
				},
			},
		},
	}
}

func getAction(name string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		context := action.Context{}
		context.Args = c.Args()
		flags := map[string]string{}
		for _, v := range c.FlagNames() {
			b := c.String(v)
			flags[v] = b
		}
		context.Flags = flags
		actionMap[name](context)
		return nil
	}
}
