package cmd

import (
	"github.com/ARGOeu/ams-push-server-cli/cmd/health"
	"github.com/spf13/cobra"
)

// NewRootCommand initialises and returns the root command of the cli
func NewRootCommand() *cobra.Command {

	var RootCmd = &cobra.Command{}
	var uri string

	// initialises the global flags
	RootCmd.PersistentFlags().StringVarP(&uri, "uri", "u", "", "-u host:port")
	RootCmd.MarkPersistentFlagRequired("uri")

	// add health check sub command
	RootCmd.AddCommand(health.NewHealthCheckCommand())

	return RootCmd
}
