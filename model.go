package logger

import "log"

// ===============================================================================
// Worker class, Worker is a log object used to log messages and Color specifies
// if colored output is to be produced
// ===============================================================================

type Worker struct {
	Minion     *log.Logger
	isColor    int
	format     string
	timeFormat string
	level      LogLevel
}

// ===============================================================================
// Info class, Contains all the info on what has to logged, time is the current time, Module is the specific module
// For which we are logging, level is the state, importance and type of message logged,
// Message contains the string to be logged, format is the format of string to be passed to sprintf
// ===============================================================================

type Info struct {
	Id       uint64
	Time     string
	Module   string
	Level    LogLevel
	Line     int
	Filename string
	Message  string
	//format   string
}

// ===============================================================================
// Logger class that is an interface to user to log messages, Module is the module for which we are testing
// worker is variable of Worker class that is used in bottom layers to log the message
// ===============================================================================

type Logger struct {
	Module string
	worker *Worker
}

type Colors struct {
	ColorMap map[LogLevel]string
}
