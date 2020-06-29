package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

var constlevel level
var logobj *logger

func ExampleSimple() {
	log := NewLogger("app")
	log.Verbose("Message A")
	log.Info("Message B")
	log.Warning("Message C")
	log.Error("Message D")
	log.Fatal("Message E")
	// Output: [VERBOSE] app 	- Message A
	// [INFO]    app 	- Message B
	// [WARNING] app 	- Message C
	// [ERROR]   app 	- Message D
	// [FATAL]   app 	- Message E
}

func ExampleFilter() {
	defaultLevel = fatallevel
	log := NewLogger("app")
	log.Verbose("Message A")
	log.Info("Message B")
	log.Warning("Message C")
	log.Error("Message D")
	log.Fatal("Message E")
	// Output: [FATAL]   app 	- Message E
}

func ExampleFilterApp() {
	defaultLevel = fatallevel
	componentsLevels["app"] = errorlevel
	log := NewLogger("app")
	log.Verbose("Message A")
	log.Info("Message B")
	log.Warning("Message C")
	log.Error("Message D")
	log.Fatal("Message E")
	// Output: [ERROR]   app 	- Message D
	// [FATAL]   app 	- Message E

}

func TestVerbose(t *testing.T) {
	//expected output
	var expectedverbose = "[VERBOSE]  	- verboselevel"
	//capture standard console output
	var buf bytes.Buffer
	logobj = new(logger)
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	//call
	logobj.Verbose("verboselevel")
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	output := buf.String()
	//test
	output = strings.Replace(output, "\n", "", 1)
	if ok := strings.Compare(output, expectedverbose); ok != 0 {
		t.Error("Expected value is 0 but got ", ok)
	}

}

func TestInfo(t *testing.T) {
	//expected output
	var expectedinfo = "[INFO]     	- infolevel"
	//capture standard console output
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	//call
	logobj.Info("infolevel")
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	output := buf.String()
	//test
	output = strings.Replace(output, "\n", "", 1)
	if ok := strings.Compare(output, expectedinfo); ok != 0 {
		t.Error("Expected value is 0 but got ", ok)
	}

}

func TestWarning(t *testing.T) {
	//expected output
	var expectedwarning = "[WARNING]  	- warninglevel"
	//capture standard console output
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	//call
	logobj.Warning("warninglevel")
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	output := buf.String()
	output = strings.Replace(output, "\n", "", 1)
	//test
	if ok := strings.Compare(output, expectedwarning); ok != 0 {
		t.Error("Expected value is 0 but got ", ok)
	}
}

func TestError(t *testing.T) {
	//expected output
	var expectederror = "[ERROR]    	- errorlevel"
	//capture standard console output
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	//call
	logobj.Error("errorlevel")
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	output := buf.String()
	output = strings.Replace(output, "\n", "", 1)
	//test
	if ok := strings.Compare(output, expectederror); ok != 0 {
		t.Error("Expected value is 0 but got ", ok)
	}

}
func TestFatal(t *testing.T) {
	//expected output
	var expectedfatal = "[FATAL]    	- fatallevel"
	//capture standard console output
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	//call
	logobj.Fatal("fatallevel")
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	output := buf.String()
	output = strings.Replace(output, "\n", "", 1)
	//test
	if ok := strings.Compare(output, expectedfatal); ok != 0 {
		t.Error("Expected value is 0 but got ", ok)
	}
}

func TestString(t *testing.T) {
	//Expected Output
	expectedfatal := "[INFO]"
	//call
	namevar := infolevel.String()
	namevar = strings.Replace(namevar, "\n", "", 0)
	namevar = strings.TrimSpace(namevar)
	//test
	if ok := strings.Compare(namevar, expectedfatal); ok != 0 {
		t.Error("Expected value is 0 but got ", ok)
	}

}

func TestNewLogger(t *testing.T) {
	lognewobj := NewLogger("Verbose")
	//Test
	if lognewobj == nil {
		t.Errorf("Expected logger object but got nil object")
	}
}

func TestPrintSettings(t *testing.T) {
	//Expected output
	expectedoutput := "Default : [VERBOSE]"
	//capture standard console output
	var buf bytes.Buffer
	logobj = new(logger)
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	//call
	defaultLevel = verboselevel
	PrintSettings()
	w.Close()
	os.Stdout = old
	io.Copy(&buf, r)
	output := buf.String()
	output = strings.Replace(output, "\n", "", 1)
	output = strings.TrimSpace(output)
	//test
	if ok := strings.Compare(output, expectedoutput); ok != 0 {
		t.Error("Got     :", []byte(output))
		t.Error("Expected:", []byte(expectedoutput))
		t.Error("Expected value is \"", expectedoutput, "\" but got \"", output, "\"")
	}
}

func TestStrtolevel(t *testing.T) {

	var fatal = "FATAL"
	var errorvar = "ERROR"
	var warning = "WARNING"
	var info = "INFO"
	var Verbose = "VERBOSE"

	//Test ERROR
	errvar, err := strtolevel(errorvar)
	if errvar != errorlevel && err != nil {
		t.Error("Expected fatal but got", errvar, err)
	}
	//Test WARNING
	war, err := strtolevel(warning)
	if war != warninglevel && err != nil {
		t.Error("Expected fatal but got", war, err)
	}
	//Test INFO
	inf, err := strtolevel(info)
	if inf != infolevel && err != nil {
		t.Error("Expected fatal but got", inf)
	}
	//Test VERBOSE
	ver, err := strtolevel(Verbose)
	if ver != verboselevel && err != nil {
		t.Error("Expected fatal but got", ver)
	}
	//Test FATAL
	fat, err := strtolevel(fatal)
	if fat != fatallevel && err != nil {
		t.Error("Expected fatal but got", fat, err)
	}

	unknown, err := strtolevel("Message")
	if unknown == ver {
		t.Errorf("Expected unknown object but got nil object")
	}

}
func init() {
	defaultLevel = verboselevel

}
