package main

import (
	"github.com/abhiyankhanal/go-logger"
	"os"
)

var (
	myColors = map[logger.LogLevel]string{
		logger.CriticalLevel: logger.ColorString(logger.Red),
		logger.ErrorLevel:    logger.ColorString(logger.Red),
		logger.WarningLevel:  logger.ColorString(logger.Yellow),
		logger.NoticeLevel:   logger.ColorString(logger.Blue),
		logger.DebugLevel:    logger.ColorString(logger.Magenta),
		logger.InfoLevel:     logger.ColorString(logger.Green),
	}
)

func main() {
	// Third option(io.Writer) is optional, defaults to os.Stderr
	// Fourth option is optional as well, else if we want custom colors then we can make by following way
	log, err := logger.New("test", 1, os.Stderr, myColors)
	if err != nil {
		panic(err) // Check for error
	}

	// Critical
	log.Critical("This is Critical!")
	// Debug
	log.Debug("This is Debug!")
	// Warning
	log.Warning("This is Warning!")
	// Error
	log.Error("This is Error!")
	// Notice
	log.Notice("This is Notice!")
	// Info
	log.Info("This is Info!")

	// Warning with formatting
	log.SetFormat("[%{level}-%{id}] %{module} %{filename} %{file}:%{line} %{message}")
	log.Warning("This is Warning!  --with customized format")
	// Set Default Format, this will only show message
	logger.SetDefaultFormat("%{message}")
	log2, _ := logger.New("pkg", 1)
	log2.Error("This is Error!")
}
