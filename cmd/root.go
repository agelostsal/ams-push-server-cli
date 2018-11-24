package cmd

import (
	"github.com/ARGOeu/ams-push-server-cli/cmd/health"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

var uri string

func init() {

	RootCmd.PersistentFlags().StringVarP(&uri, "uri", "u", "", "-u host:port")
	RootCmd.MarkPersistentFlagRequired("uri")

	RootCmd.AddCommand(health.NewHealthCheckCommand())
}
