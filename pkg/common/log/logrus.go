package log

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"open-im/pkg/common/config"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger *Logger

type Logger struct {
	*logrus.Logger
	Pid int
}

func init() {
	fmt.Println("init logger")
	logger = loggerInit("")
}

func NewPrivateLog(moduleName string) {
	logger = loggerInit(moduleName)
}

func loggerInit(moduleName string) *Logger {
	var logger = logrus.New()
	//All logs will be printed
	logger.SetLevel(logrus.Level(config.Config.Log.RemainLogLevel))
	//Close std console output
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err.Error())
	}
	writer := bufio.NewWriter(src)
	logger.SetOutput(writer)
	//logger.SetOutput(os.Stdout)
	//Log Console Print Style Setting
	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})
	//File name and line number display hook
	logger.AddHook(newFileHook())

	//Send logs to elasticsearch hook
	if config.Config.Log.ElasticSearchSwitch {
		logger.AddHook(newEsHook(moduleName))
	}
	//Log file segmentation hook
	hook := NewLfsHook(time.Duration(config.Config.Log.RotationTime)*time.Hour, config.Config.Log.RemainRotationCount, moduleName)
	logger.AddHook(hook)
	return &Logger{
		logger,
		os.Getpid(),
	}
}

func NewLfsHook(rotationTime time.Duration, maxRemainNum uint, moduleName string) logrus.Hook {
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum, "debug", moduleName),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum, "info", moduleName),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum, "warn", moduleName),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum, "error", moduleName),
	}, &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})
	return lfsHook
}

func initRotateLogs(rotationTime time.Duration, maxRemainNum uint, level string, moduleName string) *rotatelogs.RotateLogs {
	if moduleName != "" {
		moduleName = moduleName + "."
	}
	writer, err := rotatelogs.New(
		config.Config.Log.StorageLocation+moduleName+level+"."+"%Y-%m-%d",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxRemainNum),
	)
	if err != nil {
		panic(err.Error())
	} else {
		return writer
	}
}

func NewError(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Errorln(args)
}

func Info(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Infoln(args)
}

func Error(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Errorln(args)
}

func Debug(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Debugln(args)
}

func NewDebug(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Debugln(args)
}

func NewInfo(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Infoln(args)
}
