package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ashwineaso/go-asl-parser/src"
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
			log.Printf("Validating ASL file: %s", jsonPath)
			content, err := os.ReadFile(jsonPath)
			if err != nil {
				log.Fatalf("failed to read file: %v", err)
			}

			if err := src.ValidateSchema(content); err != nil {
				log.Fatalf("failed to validate schema: %v", err)
			}

			log.Println("ASL file is valid")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
