package flags

import "github.com/urfave/cli/v2"

var Email = []cli.Flag{
	&cli.StringFlag{
		Name:    "file",
		Aliases: []string{"files"},
	},
	&cli.StringFlag{
		Name:     "password",
		Required: true,
		EnvVars:  []string{"EMAIL_PASSWORD"},
	},
	&cli.StringFlag{
		Name:     "host",
		Required: true,
		EnvVars:  []string{"EMAIL_HOST"},
	},
	&cli.StringFlag{
		Name:     "from",
		Required: true,
		EnvVars:  []string{"EMAIL_FROM"},
	},
	&cli.StringFlag{
		Name:    "port",
		Value:   "587",
		EnvVars: []string{"EMAIL_PORT"},
	},
	&cli.StringFlag{
		Name:     "template",
		Required: true,
		EnvVars:  []string{"EMAIL_TEMPLATE"},
	},
	&cli.StringFlag{
		Name:     "to",
		Required: true,
		Usage:    "to send to mulitples pass them ',' seperated",
	},
	&cli.StringFlag{
		Name:     "subject",
		Required: true,
	},
	&cli.StringFlag{
		Name:  "data",
		Usage: "provide path to json file",
	},
}
