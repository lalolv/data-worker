package cmd

import (
	"github.com/gookit/config"
	"github.com/gookit/gcli/v2"
	"github.com/lalolv/data-worker/handler"
	"github.com/lalolv/goutil"
)

// options for the command
var buildOpts = struct {
	c string
}{}

// BuildCommand build command
func BuildCommand() *gcli.Command {
	cmd := &gcli.Command{
		Func: buildExecute,
		Name: "build",
		// Aliases: []string{"exp", "ex"},
		UseFor: "this is a description message",
		// {$binName} {$cmd} is help vars. '{$cmd}' will replace to 'example'
		Examples: `{$binName} {$cmd} -c config.json`,
	}

	cmd.StrOpt(&buildOpts.c, "config", "c", "value", "the config option")

	return cmd
}

// command running
func buildExecute(c *gcli.Command, args []string) error {

	// add driver for support yaml content
	// config.AddDriver(json.Driver)
	// Load json file
	err := config.LoadFiles(buildOpts.c)
	if err != nil {
		return err
	}

	// Get dict path
	dictPath, _ := config.String("dict_path")
	// Get build config
	buildInfo, _ := config.StringMap("build")
	// Get some basic params
	buildCount := goutil.Str2Int(buildInfo["count"])

	// Get all fields
	fields, _ := config.Get("fields")

	// Load dict data by fields
	handler.LoadDicts(fields.([]interface{}), float64(buildCount), dictPath)

	// Build json file
	handler.Build(buildInfo["path"], buildInfo["name"], buildCount, fields.([]interface{}))

	return nil
}
