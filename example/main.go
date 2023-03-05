package main

import "os"
import "logger"

var (
	minimalColors = map[logger.LogLevel]string{
		logger.CriticalLevel: logger.ColorString(31),
		logger.ErrorLevel:    logger.ColorString(31),
		logger.WarningLevel:  logger.ColorString(31),
		logger.NoticeLevel:   logger.ColorString(31),
		logger.DebugLevel:    logger.ColorString(31),
		logger.InfoLevel:     logger.ColorString(31),
	}
)

func main() {
	// Third option(io.Writer) is optional, defaults to os.Stderr
	log, err := logger.New("test", 1, os.Stderr, minimalColors)
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
	log2, _ := logger.New("pkg", 1, os.Stdout)
	log2.Error("This is Error!")
}
