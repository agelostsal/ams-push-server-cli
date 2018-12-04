package activate

import (
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

type subActivateOptions struct {
	FullName     string
	FullTopic    string
	PushEndpoint string
	RetryType    string
	RetryPeriod  uint32
}

// NewSubscriptionActivateCommand initialises a subscription activation command
func NewSubscriptionActivateCommand() *cobra.Command {

	var subOpts subActivateOptions

	activateSubCmd := &cobra.Command{
		Use:   "activate",
		Short: "Activates a subscription on a push server",
		Long:  "The activate command informs a push server to start the push functionality for the given subscription",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// get the global uri flag
			uri, _ := cmd.Flags().GetString("uri")
			// perform the subscription activation
			cmd.Println(grpcclient.New(uri).ActivateSubscription(subOpts.FullName, subOpts.FullTopic, subOpts.PushEndpoint, subOpts.RetryType, subOpts.RetryPeriod).Result())
		},
	}

	activateSubCmd.PersistentFlags().StringVarP(&subOpts.FullName, "full-sub", "s", "", "-s /projects/projectname/subscriptions/subanme")
	activateSubCmd.MarkPersistentFlagRequired("full-sub")

	activateSubCmd.PersistentFlags().StringVarP(&subOpts.FullTopic, "full-topic", "f", "", "-f /projects/projectname/topics/topicname")
	activateSubCmd.MarkPersistentFlagRequired("full-topic")

	activateSubCmd.PersistentFlags().StringVarP(&subOpts.PushEndpoint, "push-endpoint", "e", "", "-e https://127.0.0.1:5000/receive_here")
	activateSubCmd.MarkPersistentFlagRequired("push-endpoint")

	activateSubCmd.PersistentFlags().StringVarP(&subOpts.RetryType, "retry-type", "t", "linear", "-t linear")
	activateSubCmd.MarkPersistentFlagRequired("retry-type")

	activateSubCmd.PersistentFlags().Uint32VarP(&subOpts.RetryPeriod, "retry-period", "p", 300, "-p 300")
	activateSubCmd.MarkPersistentFlagRequired("retry-period")

	return activateSubCmd

}
