package cmd

import (
	"github.com/aamirmousavi/dong/cmd/server"
	"github.com/spf13/cobra"
)

var root = cobra.Command{}

func init() {
	root.AddCommand(
		server.ServerCmd,
	)
}

func Execute() error {
	return root.Execute()
}
