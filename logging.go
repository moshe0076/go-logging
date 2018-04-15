package logging

import (
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)

func init() {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatalln("Failed to discover host name")
	}

	Trace = log.New(
		os.Stdout,
		"TRACE "+hostName+" ",
		log.Ldate|log.Lmicroseconds|log.Llongfile)

	Info = log.New(
		os.Stdout,
		"INFO "+hostName+" ",
		log.Ldate|log.Lmicroseconds|log.Llongfile)

	Warning = log.New(
		os.Stdout,
		"WARNING "+hostName+" ",
		log.Ldate|log.Lmicroseconds|log.Llongfile)

	Error = log.New(
		os.Stderr,
		"ERROR "+hostName+" ",
		log.Ldate|log.Lmicroseconds|log.Llongfile)

	Trace.Println("Loggin Was Initialized Successfully")
}
