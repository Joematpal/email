package flags

import "github.com/urfave/cli/v2"

var Email = []cli.Flag{
	&cli.StringFlag{
		Name:    "file",
		Aliases: []string{"files"},
	},
	&cli.StringFlag{
		Name:    "password",
		EnvVars: []string{"EMAIL_PASSWORD"},
	},
	&cli.StringFlag{
		Name:    "host",
		EnvVars: []string{"EMAIL_HOST"},
	},
	&cli.StringFlag{
		Name:    "from",
		EnvVars: []string{"EMAIL_FROM"},
	},
	&cli.StringFlag{
		Name:    "port",
		Value:   "587",
		EnvVars: []string{"EMAIL_PORT"},
	},
	&cli.StringFlag{
		Name:    "template",
		EnvVars: []string{"EMAIL_TEMPLATE"},
	},
	&cli.StringFlag{
		Name: "to",
	},
	&cli.StringFlag{
		Name: "subject",
	},
	&cli.StringFlag{
		Name:  "data",
		Usage: "provide path to json file",
	},
}
