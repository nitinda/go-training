package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

// empty struct
// Signal only Channel, to recive and send message;
// Zero memory allocation
var doneCh = make(chan struct{})

func logger() {
	// for entry := range logCh {
	// fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T13:02:09"), entry.severity, entry.message)
	// }
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T13:02:09"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
