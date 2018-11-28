package deactivate

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeactivateSubscriptionCmdTestSuite struct {
	suite.Suite
}

// TestNewSubscriptionDeactivateCommand tests that an deactivate subscription command has been initialised properly
func (suite *DeactivateSubscriptionCmdTestSuite) TestNewSubscriptionDeactivateCommand() {

	ds := NewSubscriptionDeactivateCommand()
	// despite not using the output here, we save the output to a buffer so we don't pollute the std.out with the help option
	b := new(bytes.Buffer)
	ds.SetOutput(b)
	ds.SetArgs([]string{"-h"})
	ds.Execute()

	suite.Equal("deactivate", ds.Use)
	suite.Equal("Deactivates a subscription on a push server", ds.Short)
	suite.Equal("The deactivate command informs a push server to stop the push functionality for the given subscription", ds.Long)
	suite.True(ds.DisableFlagsInUseLine)

	sn := ds.Flags().Lookup("full-sub")
	suite.Equal("full-sub", sn.Name)
	suite.Equal("s", sn.Shorthand)
	suite.Equal("-s /projects/projectname/subscriptions/subanme", sn.Usage)
	suite.Equal("string", sn.Value.Type())
}

// TestSubscriptionDeactivateCmdOutput tests various outputs of the deactivate subscription command
func (suite *DeactivateSubscriptionCmdTestSuite) TestSubscriptionDeactivateCmdOutput() {

	ds := NewSubscriptionDeactivateCommand()
	b := new(bytes.Buffer)
	ds.SetOutput(b)
	ds.SetArgs([]string{"-h"})
	ds.Execute()

	expectedOut := "The deactivate command informs a push server to stop the push functionality for the given subscription\n\n" +
		"Usage:\n" +
		"  deactivate\n\n" +
		"Flags:\n" +
		"  -s, --full-sub string   -s /projects/projectname/subscriptions/subanme\n" +
		"  -h, --help              help for deactivate\n"

	// test the help output
	suite.Equal(expectedOut, b.String())

	ds2 := NewSubscriptionDeactivateCommand()
	b2 := new(bytes.Buffer)
	ds2.SetOutput(b2)
	ds2.SetArgs([]string{"deactivate",
		"--full-sub=/projects/p1/subscriptions/s1"})

	ds2.Run = func(cmd *cobra.Command, args []string) {
		// get the global uri flag
		s, _ := cmd.Flags().GetString("full-sub")
		// perform the subscription deactivation
		cmd.Printf("Success: Subscription %v deactivated\n", s)
	}

	ds2.Execute()

	// test the success output
	suite.Equal("Success: Subscription /projects/p1/subscriptions/s1 deactivated\n", b2.String())

	b2.Reset()
	ds2.Run = func(cmd *cobra.Command, args []string) {
		// get the global uri flag
		s, _ := cmd.Flags().GetString("full-sub")
		// perform the subscription deactivation
		cmd.Printf("Error: Subscription %v is not active\n", s)
	}

	ds2.Execute()

	// test the error output
	suite.Equal("Error: Subscription /projects/p1/subscriptions/s1 is not active\n", b2.String())

	ds3 := NewSubscriptionDeactivateCommand()
	b2.Reset()
	ds3.SetOutput(b2)
	ds3.SetArgs([]string{})
	ds3.Execute()
	expectedOutput2 := "Error: required flag(s) \"full-sub\" not set\n" +
		"Usage:\n" +
		"  deactivate\n\n" +
		"Flags:\n" +
		"  -s, --full-sub string   -s /projects/projectname/subscriptions/subanme\n" +
		"  -h, --help              help for deactivate\n\n"

	// test missing flags
	suite.Equal(expectedOutput2, b2.String())
}

func TestDeactivateSubscriptionCmdTestSuite(t *testing.T) {
	suite.Run(t, new(DeactivateSubscriptionCmdTestSuite))
}
