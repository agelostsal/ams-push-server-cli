package health

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HealthCmdTestSuite struct {
	suite.Suite
}

// TestNewHealthCheckCommand tests that the health check command has been initialised properly
func (suite *HealthCmdTestSuite) TestNewHealthCheckCommand() {

	hc := NewHealthCheckCommand()

	suite.Equal("health", hc.Use)
	suite.Equal("Performs a health check call", hc.Short)
	suite.Equal("The health command is used to perform a health check to the grpc server", hc.Long)
	suite.True(hc.DisableFlagsInUseLine)
}

func (suite *HealthCmdTestSuite) TestHealthCheckCmdOutput() {

	hc := NewHealthCheckCommand()
	b := new(bytes.Buffer)

	// test the output of the help option
	hc.SetArgs([]string{"-h"})
	hc.SetOutput(b)
	hc.Execute()

	expOut :=
		"The health command is used to perform a health check to the grpc server\n\n" +
			"Usage:\n" +
			"  health\n\n" +
			"Flags:\n" +
			"  -h, --help   help for health\n"

	suite.Equal(expOut, b.String())

	// test the output of the health option(success)
	hc2 := NewHealthCheckCommand()
	b2 := new(bytes.Buffer)
	hc2.SetArgs([]string{"health"})
	hc2.SetOutput(b2)
	// first override the run function in order to mock it
	hc2.Run = func(cmd *cobra.Command, args []string) {
		cmd.Println("Success: SERVING")
	}

	hc2.Execute()

	suite.Equal("Success: SERVING\n", b2.String())

	// test the output of the health option(error)
	b2.Reset()
	hc2.Run = func(cmd *cobra.Command, args []string) {
		cmd.Println("Error: NOT_SERVING")
	}

	hc2.Execute()

	suite.Equal("Error: NOT_SERVING\n", b2.String())

}

func TestHealthCommandTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCmdTestSuite))
}
