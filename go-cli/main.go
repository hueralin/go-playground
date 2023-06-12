package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	// (&cli.App{}).Run(os.Args)
	app := &cli.App{
		Name:  "Hello",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show version",
			},
		},
		// The action to execute when no subcommands are specified
		Action: func(cCtx *cli.Context) error {
			args := cCtx.Args()
			fmt.Println("hello!")
			fmt.Printf("first arg is %v\n", args.Get(0))
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
