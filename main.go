package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/5c077m4n/dotlink/linker"
	"github.com/5c077m4n/dotlink/pathfinder"
)

func main() {
	app := &cli.App{
		Name:  "dotlink",
		Usage: "Link your dotfiles into the right place",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Value:   os.ExpandEnv("${HOME}/.config/"),
				Usage:   "Set target directory",
			},
			&cli.BoolFlag{
				Name:    "replace",
				Aliases: []string{"R"},
				Value:   false,
				Usage:   "Replace current files/directories (if any exist)",
			},
		},
		Action: func(ctx *cli.Context) error {
			args := ctx.Args()
			target := ctx.String("target")
			replace := ctx.Bool("replace")

			dirs, error := pathfinder.PathFind(args)
			if error != nil {
				return error
			}

			return linker.Link(dirs, target, replace)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
