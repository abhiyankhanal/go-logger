[![goreport](https://goreportcard.com/badge/github.com/abhiyankhanal/go-logger)](https://goreportcard.com/report/github.com/abhiyankhanal/go-logger)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=abhiyankhanal_go-logger&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=abhiyankhanal_go-logger)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=abhiyankhanal_go-logger&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=abhiyankhanal_go-logger)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=abhiyankhanal_go-logger&metric=bugs)](https://sonarcloud.io/summary/new_code?id=abhiyankhanal_go-logger)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=abhiyankhanal_go-logger&metric=coverage)](https://sonarcloud.io/summary/new_code?id=abhiyankhanal_go-logger)
[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/summary/new_code?id=abhiyankhanal_go-logger)

# go-logger
The go-logger package aims to build informative logs for your go project.
Some of the features are
- Customizable colorful/non colored logs on the basis of different levels.
- Set formats to include the following:
  - Module
  - Time
  - FileName
  - File
  - Line
  - Level
  - Message
- 6 level of logs
	- Critical
	- Error
	- Warning
	- Notice
	- Debug
	- Info

# Installation
```
go get github.com/abhiyankhanal/go-logger
```

# Requirement
```
go1.18+
```

# Formatting

By default, the format of all log messages.
But, you can change the default format to one of your choosing.

After generating a logger, you can do it for the instance of the logger.
```go
log, _ := logger.New("name_here", 1)
log.SetFormat(format)
```
or for package
```go
logger.SetDefaultFormat(format)
```

# Set your custom colors
Available colors:
`Red, Yellow, Blue, Magenta, Green, White, Black`
```go
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
//first_option:package_name
//second_option: 0 for no color, 1 for colorful logs
//third_option: Writer type(io.Writer)
//fourth_option: Custom Colors
log, err := logger.New("test", 1, os.Stderr, myColors)
```
# Set Format
```go
/*
Available Formats:
	"%{id}"
	"%{time}"
	"%{module}"
	"%{filename}"
	"%{file}"
	"%{line}"
	"%{level}"
	"%{lvl}"
	"%{message}"
*/
log.SetFormat("[%{level}-%{id}] %{module} %{filename} %{file}:%{line} %{message}")
```
# Example program
[link here](https://github.com/abhiyankhanal/go-logger/blob/master/example/main.go)

# Screenshot
<img width="634" alt="image" src="https://user-images.githubusercontent.com/51784021/222954019-9e74261d-25c2-413c-b526-c66524b9db27.png">


# Contributors
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
    	<td align="center"><a href="https://github.com/abhiyankhanal"><img src="https://avatars.githubusercontent.com/u/51784021" width="70px;" alt="Abhiyan Khanal"/><br /><sub><b>Abhiyan Khanal</b></sub></a><br /></td>
		<td align="center"><a href="https://github.com/shankar524"><img src="https://avatars.githubusercontent.com/u/22054064" width="70px;" alt="Shankar Acharya"/><br /><sub><b>Shankar Acharya</b></sub></a><br /></td>
    </tr>
  </tbody>
</table>
