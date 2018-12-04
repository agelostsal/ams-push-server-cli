package activate

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ActivateSubscriptionCmdTestSuite struct {
	suite.Suite
}

// TestNewSubscriptionActivateCommand tests that an activate subscription command has been initialised properly
func (suite *ActivateSubscriptionCmdTestSuite) TestNewSubscriptionActivateCommand() {

	as := NewSubscriptionActivateCommand()
	// despite not using the output here, we save the output to a buffer so we don't pollute the std.out with the help option
	b := new(bytes.Buffer)
	as.SetOutput(b)
	as.SetArgs([]string{"-h"})
	as.Execute()

	suite.Equal("activate", as.Use)
	suite.Equal("Activates a subscription on a push server", as.Short)
	suite.Equal("The activate command informs a push server to start the push functionality for the given subscription", as.Long)
	suite.True(as.DisableFlagsInUseLine)

	sn := as.Flags().Lookup("full-sub")
	suite.Equal("full-sub", sn.Name)
	suite.Equal("s", sn.Shorthand)
	suite.Equal("-s /projects/projectname/subscriptions/subanme", sn.Usage)
	suite.Equal("string", sn.Value.Type())

	fn := as.Flags().Lookup("full-topic")
	suite.Equal("full-topic", fn.Name)
	suite.Equal("f", fn.Shorthand)
	suite.Equal("-f /projects/projectname/topics/topicname", fn.Usage)
	suite.Equal("string", fn.Value.Type())

	pe := as.Flags().Lookup("push-endpoint")
	suite.Equal("push-endpoint", pe.Name)
	suite.Equal("e", pe.Shorthand)
	suite.Equal("-e https://127.0.0.1:5000/receive_here", pe.Usage)
	suite.Equal("string", pe.Value.Type())

	rp := as.Flags().Lookup("retry-period")
	suite.Equal("retry-period", rp.Name)
	suite.Equal("p", rp.Shorthand)
	suite.Equal("-p 300", rp.Usage)
	suite.Equal("300", rp.DefValue)
	suite.Equal("uint32", rp.Value.Type())

	rt := as.Flags().Lookup("retry-type")
	suite.Equal("retry-type", rt.Name)
	suite.Equal("t", rt.Shorthand)
	suite.Equal("-t linear", rt.Usage)
	suite.Equal("linear", rt.DefValue)
	suite.Equal("string", rt.Value.Type())
}

// TestSubscriptionActivateCmdOutput tests various outputs of the activate subscription command
func (suite *ActivateSubscriptionCmdTestSuite) TestSubscriptionActivateCmdOutput() {

	as := NewSubscriptionActivateCommand()
	b := new(bytes.Buffer)
	as.SetOutput(b)
	as.SetArgs([]string{"-h"})
	as.Execute()

	expectedOut := "The activate command informs a push server to start the push functionality for the given subscription\n\n" +
		"Usage:\n" +
		"  activate\n\n" +
		"Flags:\n" +
		"  -s, --full-sub string        -s /projects/projectname/subscriptions/subanme\n" +
		"  -f, --full-topic string      -f /projects/projectname/topics/topicname\n" +
		"  -h, --help                   help for activate\n" +
		"  -e, --push-endpoint string   -e https://127.0.0.1:5000/receive_here\n" +
		"  -p, --retry-period uint32    -p 300 (default 300)\n" +
		"  -t, --retry-type string      -t linear (default \"linear\")\n"

	// test the help output
	suite.Equal(expectedOut, b.String())

	as2 := NewSubscriptionActivateCommand()
	b2 := new(bytes.Buffer)
	as2.SetOutput(b2)
	as2.SetArgs([]string{"activate",
		"--full-sub=/projects/p1/subscriptions/s1",
		"--full-topic=/projects/p1/topics/t1",
		"--push-endpoint=https:example.com:5000/receive_here",
		"--retry-period=300",
		"--retry-type=linear",
	})
	as2.Run = func(cmd *cobra.Command, args []string) {
		// get the global uri flag
		s, _ := cmd.Flags().GetString("full-sub")
		// perform the subscription activation
		cmd.Printf("Success: Subscription %v activated\n", s)
	}

	as2.Execute()

	// test the success output
	suite.Equal("Success: Subscription /projects/p1/subscriptions/s1 activated\n", b2.String())

	b2.Reset()
	as2.Run = func(cmd *cobra.Command, args []string) {
		// get the global uri flag
		s, _ := cmd.Flags().GetString("full-sub")
		// perform the subscription activation
		cmd.Printf("Error: Subscription %v is already activated\n", s)
	}

	as2.Execute()

	// test the error output
	suite.Equal("Error: Subscription /projects/p1/subscriptions/s1 is already activated\n", b2.String())

	as3 := NewSubscriptionActivateCommand()
	b2.Reset()
	as3.SetOutput(b2)
	as3.SetArgs([]string{})
	as3.Execute()
	expectedOutput2 := "Error: required flag(s) \"full-sub\", \"full-topic\", \"push-endpoint\", \"retry-period\", \"retry-type\" not set\n" +
		"Usage:\n" +
		"  activate\n\n" +
		"Flags:\n" +
		"  -s, --full-sub string        -s /projects/projectname/subscriptions/subanme\n" +
		"  -f, --full-topic string      -f /projects/projectname/topics/topicname\n" +
		"  -h, --help                   help for activate\n" +
		"  -e, --push-endpoint string   -e https://127.0.0.1:5000/receive_here\n" +
		"  -p, --retry-period uint32    -p 300 (default 300)\n" +
		"  -t, --retry-type string      -t linear (default \"linear\")\n\n"

	// test missing flags
	suite.Equal(expectedOutput2, b2.String())
}

func TestNewSubscriptionActivateCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ActivateSubscriptionCmdTestSuite))
}
