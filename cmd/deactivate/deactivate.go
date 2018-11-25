package deactivate

import (
	"fmt"
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

var subName string

func NewSubscriptionDeactivateCommand() *cobra.Command {

	deactivateSubCmd := &cobra.Command{
		Use: "deactivate",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			t1, err := cmd.Flags().GetString("urid")
			fmt.Printf("%+v", err)
			grpcclient.New(t1).DeactivateSubscription(subName).Result()
		},
	}

	return deactivateSubCmd

}
