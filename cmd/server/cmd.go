package server

import (
	"fmt"
	"os"

	"github.com/aamirmousavi/dong/utils/sms"
	"github.com/spf13/cobra"
)

var (
	addr        = ":8080"
	mongodbAddr = "mongodb://192.168.1.3:27017"
	smsUsername = ""
	smsPassword = ""
)

func init() {
	ServerCmd.Flags().StringVarP(&addr, "addr", "a", addr, "server address")
	ServerCmd.Flags().StringVarP(&mongodbAddr, "mongodb", "m", mongodbAddr, "mongodb address")
	ServerCmd.Flags().StringVarP(&smsUsername, "sms-username", "u", smsUsername, "sms username")
	ServerCmd.Flags().StringVarP(&smsPassword, "sms-password", "p", smsPassword, "sms password")
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server command",
	Run: func(cmd *cobra.Command, args []string) {

		if smsUsername != "" && smsPassword != "" {
			sms.InitSetUsernameAndPassword(smsUsername, smsPassword)
		}

		if err := run(mongodbAddr, addr); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}
