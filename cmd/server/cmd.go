package server

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	addr        = ":8080"
	mongodbAddr = "mongodb://localhost:27017"
)

func init() {
	ServerCmd.Flags().StringVarP(&addr, "addr", "a", addr, "server address")
	ServerCmd.Flags().StringVarP(&mongodbAddr, "mongodb", "m", mongodbAddr, "mongodb address")
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server command",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(mongodbAddr, addr); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}
