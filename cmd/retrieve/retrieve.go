package retrieve

import (
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

// NewSubscriptionListOneCommand initialises a subscription retrieval command
func NewSubscriptionListOneCommand() *cobra.Command {

	var subName string

	retrieveSubCmd := &cobra.Command{
		Use:   "list-one",
		Short: "Retrieves information about a subscription currently active on the push server",
		Long:  "The list-one command retrieves all information related to the given subscription name from the push server",
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
