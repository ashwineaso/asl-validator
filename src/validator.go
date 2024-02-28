package src

import (
	"embed"
	_ "embed"
	"encoding/json"
	"io/fs"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schemas/*.json
var schemaFolder embed.FS

func ValidateSchema(payload []byte) error {
	c := jsonschema.NewCompiler()
	c.Draft = jsonschema.Draft7

	// Add all schema in the folder to the compiler
	if err := fs.WalkDir(schemaFolder, "schemas", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		var schema []byte
		schema, err = schemaFolder.ReadFile(path)
		if err != nil {
			return err
		}

		err = c.AddResource(path, strings.NewReader(string(schema)))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	// Compile all schemas
	sch, err := c.Compile("schemas/state-machine.json")
	if err != nil {
		return err
	}

	var v interface{}
	if err = json.Unmarshal(payload, &v); err != nil {
		return err
	}

	if err = sch.Validate(v); err != nil {
		return err
	}

	return nil
}
