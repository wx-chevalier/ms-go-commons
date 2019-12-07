package fileutils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Basic imports

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	var filename string

	f, err := ioutil.TempFile("", "messages_test")
	require.Nil(suite.T(), err)
	filename = f.Name()
	require.Contains(suite.T(), filename, "messages_test")

	input := &types.Struct{
		Fields: map[string]*types.Value{
			"foo": {
				Kind: &types.Value_StringValue{StringValue: "bar"},
			},
		},
	}

	err = WriteToFile(filename, input)
	require.Nil(suite.T(), err)

	b, err := ioutil.ReadFile(filename)
	require.Nil(suite.T(), err)

	require.Equal(suite.T(), string(b), "foo: bar\n")

	var output types.Struct
	err = ReadFileInto(filename, &output)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), output, *input)

	os.RemoveAll(filename)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
