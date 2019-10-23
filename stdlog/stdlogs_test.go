package stdlog

import (
	"log"
	"testing"
)

func TestStdLogDefault(t *testing.T) {
	testErrlogger := log.New(Err, "errlog:", log.Ldate)
	testOutlogger := log.New(Out, "outlog:", log.Ldate)

	testErrlogger.Println("writing to ", Err.Filename)
	testOutlogger.Println("writing to ", Out.Filename)
	t.Logf("writing to %v", Err.Filename)
	t.Logf("writing to %v", Out.Filename)
}

func TestStdLogKnown(t *testing.T) {
	SetLogFiles("../rpn.log", "../rpn_error.log")

	testErrlogger := log.New(Err, "errlog:", log.Ldate)
	testOutlogger := log.New(Out, "outlog:", log.Ldate)

	testErrlogger.Println("writing to ", Err.Filename)
	testOutlogger.Println("writing to ", Out.Filename)
	t.Log("writing to ", Err.Filename)
	t.Log("writing to ", Out.Filename)
}
