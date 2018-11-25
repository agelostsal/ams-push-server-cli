package cmd

import "github.com/spf13/cobra"

// NewRootCommand initialises and returns the root command of the cli
func NewRootCommand() *cobra.Command {

	var RootCmd = &cobra.Command{}
	var uri string

	RootCmd.PersistentFlags().StringVarP(&uri, "uri", "u", "", "-u host:port")
	RootCmd.MarkPersistentFlagRequired("uri")
	return RootCmd
}
