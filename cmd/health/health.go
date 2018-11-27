package health

import (
	"github.com/ARGOeu/ams-push-server-cli/grpc/client"
	"github.com/spf13/cobra"
)

// NewHealthCheckCommand initialises the heath check command
func NewHealthCheckCommand() *cobra.Command {

	healthCheckCmd := &cobra.Command{
		Use:   "health",
		Short: "Performs a health check call",
		Long:  "The health command is used to perform a health check to the grpc server",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// get the global uri flag
			uri, _ := cmd.Flags().GetString("uri")
			// perform the health check
			cmd.Println(grpcclient.New(uri).HealthCheck().Result())
		},
	}

	return healthCheckCmd
}
