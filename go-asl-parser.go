package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	var jsonPath string

	app := &cli.App{
		Name:  "go-asl-validator",
		Usage: "validate the ASL file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "json-path",
				Usage:       "json path to the ASL file `FILE`",
				Destination: &jsonPath,
			},
		},
		Action: func(*cli.Context) error {
			//TODO: validate the ASL file
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
