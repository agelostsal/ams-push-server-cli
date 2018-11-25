package activate

import (
	"fmt"
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

type SubActivateOptions struct {
	FullName     string
	FullTopic    string
	PushEndpoint string
	RetryType    string
	RetryPeriod  uint32
}

var subOpts SubActivateOptions

func NewSubscriptionActivateCommand() *cobra.Command {

	activateSubCmd := &cobra.Command{
		Use:                   "activate ([-f FILENAME] | TYPE [(NAME | -l label | --all)])",
		Short:                 "llll",
		Long:                  "ffff",
		Example:               "fgsgvaav",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			t1, err := cmd.Flags().GetString("urid")
			fmt.Printf("%+v", err)
			grpcclient.New(t1).ActivateSubscription(subOpts.FullName, "", "", "", 90).Result()
		},
	}

	activateSubCmd.PersistentFlags().StringVarP(&subOpts.FullName, "fullname", "n", "", "-n /projects/project/subscriptions/subanme")
	activateSubCmd.MarkPersistentFlagRequired("fullname")
	activateSubCmd.PersistentFlags().StringVarP(&subOpts.FullTopic, "fulltopic", "t", "", "-t /projects/project/topics/topic")
	activateSubCmd.MarkPersistentFlagRequired("fulltopic")
	activateSubCmd.PersistentFlags().StringVarP(&subOpts.PushEndpoint, "pushendpoint", "e", "", "-e https://127.0.0.1:5000/receive_here")
	activateSubCmd.MarkPersistentFlagRequired("pushendpoint")
	activateSubCmd.PersistentFlags().StringVarP(&subOpts.RetryType, "retrytype", "r", "", "-r linear")
	activateSubCmd.MarkPersistentFlagRequired("retrytype")
	activateSubCmd.PersistentFlags().Uint32VarP(&subOpts.RetryPeriod, "retryperiod", "p", 0, "-p 300")
	activateSubCmd.MarkPersistentFlagRequired("retryperiod")

	return activateSubCmd

}
