package list

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fatih/color"
	"proctor/daemon"
	"proctor/io"
	proc_metadata "proctor/proctord/jobs/metadata"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListCmdTestSuite struct {
	suite.Suite
	mockPrinter        *io.MockPrinter
	mockProctorDClient *daemon.MockClient
	testListCmd        *cobra.Command
}

func (s *ListCmdTestSuite) SetupTest() {
	s.mockPrinter = &io.MockPrinter{}
	s.mockProctorDClient = &daemon.MockClient{}
	s.testListCmd = NewCmd(s.mockPrinter, s.mockProctorDClient)
}

func (s *ListCmdTestSuite) TestListCmdUsage() {
	assert.Equal(s.T(), "list", s.testListCmd.Use)
}

func (s *ListCmdTestSuite) TestListCmdHelp() {
	assert.Equal(s.T(), "List procs available for execution", s.testListCmd.Short)
	assert.Equal(s.T(), "List procs available for execution", s.testListCmd.Long)
	assert.Equal(s.T(), "proctor list", s.testListCmd.Example)
}

func (s *ListCmdTestSuite) TestListCmdRun() {
	procOne := proc_metadata.Metadata{
		Name:        "one",
		Description: "proc one description",
	}
	procTwo := proc_metadata.Metadata{
		Name:        "two",
		Description: "proc two description",
	}
	procList := []proc_metadata.Metadata{procOne, procTwo}

	s.mockProctorDClient.On("ListProcs").Return(procList, nil).Once()

	s.mockPrinter.On("Println", "List of Procs:\n", color.FgGreen).Once()
	s.mockPrinter.On("Println", fmt.Sprintf("%-40s %-100s", procOne.Name, procOne.Description), color.Reset).Once()
	s.mockPrinter.On("Println", fmt.Sprintf("%-40s %-100s", procTwo.Name, procTwo.Description), color.Reset).Once()
	s.mockPrinter.On("Println", "\nFor detailed information of any proc, run:\nproctor describe <proc_name>", color.FgGreen).Once()
	s.testListCmd.Run(&cobra.Command{}, []string{})

	s.mockProctorDClient.AssertExpectations(s.T())
	s.mockPrinter.AssertExpectations(s.T())
}

func (s *ListCmdTestSuite) TestListCmdRunProctorDClientFailure() {
	s.mockProctorDClient.On("ListProcs").Return([]proc_metadata.Metadata{}, errors.New("Error!!!\nUnknown Error.")).Once()
	s.mockPrinter.On("Println", "Error!!!\nUnknown Error.", color.FgRed).Once()

	s.testListCmd.Run(&cobra.Command{}, []string{})

	s.mockProctorDClient.AssertExpectations(s.T())
	s.mockPrinter.AssertExpectations(s.T())
}

func TestListCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ListCmdTestSuite))
}
