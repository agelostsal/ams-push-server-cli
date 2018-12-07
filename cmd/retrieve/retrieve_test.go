package retrieve

import (
	"bytes"
	"github.com/ARGOeu/ams-push-server-cli/grpc/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RetrieveSubscriptionCmdTestSuite struct {
	suite.Suite
}

// TestNewSubscriptionListOneCommand tests that a list one subscription command has been initialised properly
func (suite *RetrieveSubscriptionCmdTestSuite) TestNewSubscriptionListOneCommand() {

	ds := NewSubscriptionListOneCommand()
	// despite not using the output here, we save the output to a buffer so we don't pollute the std.out with the help option
	b := new(bytes.Buffer)
	ds.SetOutput(b)
	ds.SetArgs([]string{"-h"})
	ds.Execute()

	suite.Equal("list-one", ds.Use)
	suite.Equal("Retrieves information about a subscription currently active on the push server", ds.Short)
	suite.Equal("The list-one command retrieves all information related to the given subscription name from the push server", ds.Long)
	suite.True(ds.DisableFlagsInUseLine)

	sn := ds.PersistentFlags().Lookup("full-sub")
	suite.Equal("full-sub", sn.Name)
	suite.Equal("s", sn.Shorthand)
	suite.Equal("-s /projects/projectname/subscriptions/subanme", sn.Usage)
	suite.Equal("string", sn.Value.Type())
}

// TestSubscriptionListOneCmdOutput tests various outputs of the list one subscription command
func (suite *RetrieveSubscriptionCmdTestSuite) TestSubscriptionListOneCmdOutput() {

	ds := NewSubscriptionListOneCommand()
	b := new(bytes.Buffer)
	ds.SetOutput(b)
	ds.SetArgs([]string{"-h"})
	ds.Execute()

	expectedOut := "The list-one command retrieves all information related to the given subscription name from the push server\n\n" +
		"Usage:\n" +
		"  list-one\n\n" +
		"Flags:\n" +
		"  -s, --full-sub string   -s /projects/projectname/subscriptions/subanme\n" +
		"  -h, --help              help for list-one\n"

	// test the help output
	suite.Equal(expectedOut, b.String())

	ds2 := NewSubscriptionListOneCommand()
	b2 := new(bytes.Buffer)
	ds2.SetOutput(b2)
	ds2.SetArgs([]string{"list-one",
		"--full-sub=/projects/p1/subscriptions/s1"})

	retry := ams.RetryPolicy{
		Type:   "linear",
		Period: 30,
	}

	pCfg := ams.PushConfig{
		PushEndpoint: "https://127.0.0.1:5000/receive_here",
		RetryPolicy:  &retry,
	}

	sub := ams.Subscription{
		FullName:   "projects/p1/subscription/sub1",
		FullTopic:  "projects/p1/topics/topic1",
		PushConfig: &pCfg,
	}

	getSubR := &ams.GetSubscriptionResponse{Subscription: &sub}

	ds2.Run = func(cmd *cobra.Command, args []string) {
		// get the global uri flag
		s, _ := cmd.PersistentFlags().GetString("full-sub")
		// perform the subscription deactivation
		getSubR.Subscription.FullName = s
		cmd.Printf("Success: %v\n", getSubR.String())
	}

	ds2.Execute()

	// test the success output
	suite.Equal("Success: subscription:<full_name:\"/projects/p1/subscriptions/s1\" full_topic:\"projects/p1/topics/topic1\" PushConfig:<push_endpoint:\"https://127.0.0.1:5000/receive_here\" RetryPolicy:<type:\"linear\" period:30 > > > \n", b2.String())

	b2.Reset()
	ds2.Run = func(cmd *cobra.Command, args []string) {
		// get the global uri flag
		s, _ := cmd.PersistentFlags().GetString("full-sub")
		// perform the subscription deactivation
		cmd.Printf("Error: Subscription %v is not active\n", s)
	}

	ds2.Execute()

	// test the error output
	suite.Equal("Error: Subscription /projects/p1/subscriptions/s1 is not active\n", b2.String())

	ds3 := NewSubscriptionListOneCommand()
	b2.Reset()
	ds3.SetOutput(b2)
	ds3.SetArgs([]string{})
	ds3.Execute()
	expectedOutput2 := "Error: required flag(s) \"full-sub\" not set\n" +
		"Usage:\n" +
		"  list-one\n\n" +
		"Flags:\n" +
		"  -s, --full-sub string   -s /projects/projectname/subscriptions/subanme\n" +
		"  -h, --help              help for list-one\n\n"

	// test missing flags
	suite.Equal(expectedOutput2, b2.String())
}

// TestNewSubscriptionListManyCommand tests that a list many subscription command has been initialised properly
func (suite *RetrieveSubscriptionCmdTestSuite) TestNewSubscriptionListManyCommand() {

	ds := NewSubscriptionListManyCommand()
	// despite not using the output here, we save the output to a buffer so we don't pollute the std.out with the help option
	b := new(bytes.Buffer)
	ds.SetOutput(b)
	ds.SetArgs([]string{"-h"})
	ds.Execute()

	suite.Equal("list-many", ds.Use)
	suite.Equal("Retrieves all currently active subscriptions", ds.Short)
	suite.Equal("The list-many command retrieves the names of all currently active subscriptions on the push server", ds.Long)
	suite.True(ds.DisableFlagsInUseLine)
}

// TestSubscriptionListManyCmdOutput tests various outputs of the list many subscription command
func (suite *RetrieveSubscriptionCmdTestSuite) TestSubscriptionListManyCmdOutput() {

	ds := NewSubscriptionListManyCommand()
	b := new(bytes.Buffer)
	ds.SetOutput(b)
	ds.SetArgs([]string{"-h"})
	ds.Execute()

	expectedOut := "The list-many command retrieves the names of all currently active subscriptions on the push server\n\n" +
		"Usage:\n" +
		"  list-many\n\n" +
		"Flags:\n" +
		"  -h, --help   help for list-many\n"

		// test the help output
	suite.Equal(expectedOut, b.String())

	ds2 := NewSubscriptionListManyCommand()
	b.Reset()
	ds2.SetOutput(b)
	ds2.SetArgs([]string{"list-many"})
	ds2.Run = func(cmd *cobra.Command, args []string) {
		cmd.Printf("Success: %v", []string{"s1", "s2", "s3"})
	}
	ds2.Execute()

	expectedOut2 := "Success: [s1 s2 s3]"

	// test the help output
	suite.Equal(expectedOut2, b.String())

}
func TestDeactivateSubscriptionCmdTestSuite(t *testing.T) {
	suite.Run(t, new(RetrieveSubscriptionCmdTestSuite))
}
