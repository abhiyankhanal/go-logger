package main

import (
	"github.com/abhiyankhanal/go-logger"
	"os"
)

var (
	myColors = map[go_logger.LogLevel]string{
		go_logger.CriticalLevel: go_logger.ColorString(go_logger.Red),
		go_logger.ErrorLevel:    go_logger.ColorString(go_logger.Red),
		go_logger.WarningLevel:  go_logger.ColorString(go_logger.Yellow),
		go_logger.NoticeLevel:   go_logger.ColorString(go_logger.Blue),
		go_logger.DebugLevel:    go_logger.ColorString(go_logger.Magenta),
		go_logger.InfoLevel:     go_logger.ColorString(go_logger.Green),
	}
)

func main() {
	// Third option(io.Writer) is optional, defaults to os.Stderr
	// Fourth option is optional as well, else if we want custom colors then we can make by following way
	log, err := go_logger.New("test", 1, os.Stderr, myColors)
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
	go_logger.SetDefaultFormat("%{message}")
	log2, _ := go_logger.New("pkg", 1)
	log2.Error("This is Error!")
}
