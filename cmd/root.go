package cmd

import (
	"github.com/ARGOeu/ams-push-server-cli/cmd/activate"
	"github.com/ARGOeu/ams-push-server-cli/cmd/deactivate"
	"github.com/ARGOeu/ams-push-server-cli/cmd/health"
	"github.com/ARGOeu/ams-push-server-cli/cmd/retrieve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRootCommand initialises and returns the root command of the cli
func NewRootCommand() *cobra.Command {

	var rootCmd = &cobra.Command{}
	var uri string

	// initialises the global flags
	rootCmd.PersistentFlags().StringVarP(&uri, "uri", "u", "", "-u host:port")
	rootCmd.MarkPersistentFlagRequired("uri")

	// set the flag retrieved from viper
	// if the cli argumernt is given, it will over it
	if viper.GetString("uri") != "" {
		rootCmd.PersistentFlags().Set("uri", viper.GetString("uri"))
	}

	// add health check sub command
	rootCmd.AddCommand(health.NewHealthCheckCommand())

	// add deactivate subscription command
	rootCmd.AddCommand(deactivate.NewSubscriptionDeactivateCommand())

	// add list-one subscription command
	rootCmd.AddCommand(retrieve.NewSubscriptionListOneCommand())

	// add activate subscription sub command
	rootCmd.AddCommand(activate.NewSubscriptionActivateCommand())

	return rootCmd
}
