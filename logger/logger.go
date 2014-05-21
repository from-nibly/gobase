package logger

import (
	"io"
	"log"
	"os"
)

var (
	trace       *log.Logger
	info        *log.Logger
	warning     *log.Logger
	eRror       *log.Logger
	debug       *log.Logger
	initialized bool = false
)

func Init(traceHandle io.Writer, infoHandle io.Writer,
	warningHandle io.Writer, errorHandle io.Writer, debugHandle io.Writer) {

	initialized = true
	trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	eRror = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	debug = log.New(debugHandle,
		"DEBUG:",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Error() *log.Logger {
	if initialized {
		return eRror
	} else {
		log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile).Println("Uninitialized Logger")
		return nil
	}
}

func Warning() *log.Logger {
	if initialized {
		return warning
	} else {
		log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile).Println("Uninitialized Logger")
		return nil
	}
}
func Trace() *log.Logger {
	if initialized {
		return trace
	} else {
		log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile).Println("Uninitialized Logger")
		return nil
	}
}
func Info() *log.Logger {
	if initialized {
		return info
	} else {
		log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile).Println("Uninitialized Logger")
		return nil
	}
}

func Debug() *log.Logger {
	if initialized {
		return debug
	} else {
		log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile).Println("Uninitialized Logger")
		return nil
	}
}
