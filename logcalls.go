// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package tracelog : logcalls.go provides formatting functions.
package tracelog

import "fmt"

//** STARTED AND COMPLETED

// Started uses the Serialize destination and adds a Started tag to the log line
func Started() {
	logger.Trace.Output(2, "Started\n")
}

// Startedf uses the Serialize destination and writes a Started tag to the log line
func Startedf(format string, a ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf("Started : %s\n", fmt.Sprintf(format, a...)))
}

// Completed uses the Serialize destination and writes a Completed tag to the log line
func Completed() {
	logger.Trace.Output(2, fmt.Sprintf("Completed\n"))
}

// Completedf uses the Serialize destination and writes a Completed tag to the log line
func Completedf(format string, a ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf("Completed : %s\n", fmt.Sprintf(format, a...)))
}

// CompletedError uses the Error destination and writes a Completed tag to the log line
func CompletedError(err error) {
	logger.Error.Output(2, fmt.Sprintf("Completed : ERROR : %s\n", err))
}

// CompletedErrorf uses the Error destination and writes a Completed tag to the log line
func CompletedErrorf(err error, format string, a ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("Completed : ERROR : %s : %s\n", fmt.Sprintf(format, a...), err))
}

//** TRACE

// Trace writes to the Trace destination
func Trace(format string, a ...interface{}) {
	logger.Trace.Output(2, fmt.Sprintf(" %s\n", fmt.Sprintf(format, a...)))
}

//** INFO

// Info writes to the Info destination
func Info(format string, a ...interface{}) {
	logger.Info.Output(2, fmt.Sprintf(" %s\n", fmt.Sprintf(format, a...)))
}

//** WARNING

// Warning writes to the Warning destination
func Warning(format string, a ...interface{}) {
	logger.Warning.Output(2, fmt.Sprintf(" %s\n", fmt.Sprintf(format, a...)))
}

//** ERROR

// Error writes to the Error destination and accepts an err
func Error(err error) {
	logger.Error.Output(2, fmt.Sprintf("ERROR : %s\n", err))
}

// Errorf writes to the Error destination and accepts an err
func Errorf(err error, format string, a ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("ERROR : %s : %s\n", fmt.Sprintf(format, a...), err))
}

//** ALERT

// Alert write to the Error destination and sends email alert
func Alert(subject string, format string, a ...interface{}) {
	message := fmt.Sprintf("ALERT : %s\n", fmt.Sprintf(format, a...))
	logger.Error.Output(2, message)
	SendEmailException(subject, message)
}

// CompletedAlert write to the Error destination, writes a Completed tag to the log line and sends email alert
func CompletedAlert(subject string, format string, a ...interface{}) {
	message := fmt.Sprintf("Completed : ALERT : %s\n", fmt.Sprintf(format, a...))
	logger.Error.Output(2, message)
	SendEmailException(subject, message)
}
