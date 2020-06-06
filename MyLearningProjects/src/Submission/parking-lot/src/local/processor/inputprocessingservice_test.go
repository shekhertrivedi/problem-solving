package processor

import (
	"testing"
)

var testCommandProc ICommandProcessor

func TestProcessCommand1(t *testing.T) {
	testCommandProc = &CommandProcessorImpl{}
	testCommandProc.SetPartnerServiceInstance()
	err := testCommandProc.ProcessCommand("")
	if err == nil {
		t.Errorf("Error occured")
	}
}

func TestProcessCommand2(t *testing.T) {
	testCommandProc = &CommandProcessorImpl{}
	testCommandProc.SetPartnerServiceInstance()
	err := testCommandProc.ProcessCommand("xyz")
	if err == nil {
		t.Errorf("Error occured")
	}
}

func TestReadInputFile1(t *testing.T) {
	testCommandProc = &CommandProcessorImpl{}
	testCommandProc.SetPartnerServiceInstance()
	err := testCommandProc.ReadInputFile("")
	if err == nil {
		t.Errorf("Error occured")
	}
}

func TestReadInputFile2(t *testing.T) {
	testCommandProc = &CommandProcessorImpl{}
	testCommandProc.SetPartnerServiceInstance()
	err := testCommandProc.ReadInputFile("xyz")
	if err == nil {
		t.Errorf("Error occured")
	}
}
