package cmd

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RootCmdTestSuite struct {
	suite.Suite
}

// TestNewRootCmd tests that the root command is initialised properly
func (suite *RootCmdTestSuite) TestNewRootCmd() {

	// initialise viper
	config := `{
	"uri": "host:5555"
}`
	viper.SetConfigType("json")
	viper.ReadConfig(bytes.NewBuffer([]byte(config)))

	rtCmd := NewRootCommand()
	rtCmd.Execute()

	/// check that the flag has been pre assigned from the viper config
	u, err := rtCmd.PersistentFlags().GetString("uri")
	suite.Equal("host:5555", u)
	suite.Nil(err)

	uriFlag := rtCmd.PersistentFlags().Lookup("uri")

	suite.Equal("uri", uriFlag.Name)
	suite.Equal("u", uriFlag.Shorthand)
	suite.Equal("-u host:port", uriFlag.Usage)
	suite.Equal("string", uriFlag.Value.Type())
}

// TestRootCmdOutput tests the output
func (suite *RootCmdTestSuite) TestRootCmdOutput() {

	rtCmd := NewRootCommand()
	b := new(bytes.Buffer)

	rtCmd.SetArgs([]string{"-h"})
	rtCmd.SetOutput(b)
	rtCmd.Execute()

	expectedOutput :=
		"Usage:\n" +
			"   [command]\n\n" +
			"Available Commands:\n" +
			"  activate    Activates a subscription on a push server\n" +
			"  deactivate  Deactivates a subscription on a push server\n" +
			"  health      Performs a health check call\n" +
			"  help        Help about any command\n" +
			"  list-one    Retrieves information about a subscription currently active on the push server\n\n" +
			"Flags:\n" +
			"  -h, --help         help for this command\n" +
			"  -u, --uri string   -u host:port\n\n" +
			"Use \" [command] --help\" for more information about a command.\n"

	suite.Equal(expectedOutput, b.String())
}

// TestRootCmdFlagValues tests that the global flags set at the root command are parsed and accessed correctly
func (suite *RootCmdTestSuite) TestRootCmdFlagValues() {

	// initialise viper
	config := `{
	"uri": "host:5555"
}`
	viper.SetConfigType("json")
	viper.ReadConfig(bytes.NewBuffer([]byte(config)))

	// make sure that since the uri argument has been set through the cli, it will override viper's configuration value

	rtCmd := NewRootCommand()
	// despite not using the output here, we save the output to a buffer so we don't pollute the std.out with the help option
	b := new(bytes.Buffer)
	rtCmd.SetOutput(b)

	rtCmd.SetArgs([]string{"--uri=localhost:5555"})
	rtCmd.Execute()

	u, err := rtCmd.PersistentFlags().GetString("uri")

	rtCmd2 := NewRootCommand()
	rtCmd2.SetOutput(b)
	rtCmd2.SetArgs([]string{"-u=localhost:5555"})
	rtCmd2.Execute()

	u2, err2 := rtCmd2.PersistentFlags().GetString("uri")

	suite.Equal("localhost:5555", u)
	suite.Nil(err)

	suite.Equal("localhost:5555", u2)
	suite.Nil(err2)
}

func TestRootCmdTestSuite(t *testing.T) {
	suite.Run(t, new(RootCmdTestSuite))
}
