package deactivate

import (
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

// NewSubscriptionDeactivateCommand initialises a subscription deactivation command
func NewSubscriptionDeactivateCommand() *cobra.Command {

	var subName string

	deactivateSubCmd := &cobra.Command{
		Use:   "deactivate",
		Short: "Deactivates a subscription on a push server",
		Long:  "The deactivate command informs a push server to stop the push functionality for the given subscription",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// get the global uri flag
			uri, _ := cmd.Flags().GetString("uri")
			// perform the subscription deactivation
			cmd.Println(grpcclient.New(uri).DeactivateSubscription(subName).Result())
		},
	}

	deactivateSubCmd.PersistentFlags().StringVarP(&subName, "full-sub", "s", "", "-s /projects/project/subscriptions/subanme")
	deactivateSubCmd.MarkPersistentFlagRequired("full-sub")

	return deactivateSubCmd
}
