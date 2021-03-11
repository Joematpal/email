package main

import (
	"os"

	"github.com/joematpal/email/cmd"
)

func main() {
	if err := cmd.NewApp().Run(os.Args); err != nil {
		panic(err)
	}
}
