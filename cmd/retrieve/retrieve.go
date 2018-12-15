package retrieve

import (
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

// NewSubscriptionGetCommand initialises a subscription retrieval command
func NewSubscriptionGetCommand() *cobra.Command {

	var subName string

	retrieveSubCmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves information about a subscription currently active on the push server",
		Long:  "The get command retrieves all information related to the given subscription name from the push server",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// get the global uri flag
			uri, _ := cmd.Flags().GetString("uri")
			// perform the subscription retrieval
			cmd.Println(grpcclient.New(uri).GetSubscription(subName).Result())
		},
	}

	retrieveSubCmd.PersistentFlags().StringVarP(&subName, "full-sub", "s", "", "-s /projects/projectname/subscriptions/subanme")
	retrieveSubCmd.MarkPersistentFlagRequired("full-sub")

	return retrieveSubCmd
}

// NewSubscriptionGetAllCommand initialises a subscriptions retrieval command
func NewSubscriptionGetAllCommand() *cobra.Command {

	retrieveSubCmd := &cobra.Command{
		Use:   "get-all",
		Short: "Retrieves all currently active subscriptions",
		Long:  "The get-all command retrieves the names of all currently active subscriptions on the push server",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// get the global uri flag
			uri, _ := cmd.Flags().GetString("uri")
			// perform the subscription retrieval
			cmd.Println(grpcclient.New(uri).ListSubscriptions().Result())
		},
	}

	return retrieveSubCmd
}
