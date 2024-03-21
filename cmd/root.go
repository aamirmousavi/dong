package cmd

import "github.com/spf13/cobra"

var root = cobra.Command{}

func Execute() error {
	return root.Execute()
}
