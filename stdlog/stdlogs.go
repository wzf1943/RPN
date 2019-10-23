package stdlog

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Err = &lumberjack.Logger{
	MaxSize: 500,
	MaxAge:  28,
}

var Out = &lumberjack.Logger{
	MaxSize: 500,
	MaxAge:  28,
}

func SetLogFiles(outfile, errfile string) {
	log.Println("stdlog.Err bound to rotating log:", errfile)
	Err.Filename = errfile
	Err.Close()
	log.Println("stdlog.Out bound to rotating log:", outfile)
	Out.Filename = outfile
	Out.Close()
}
