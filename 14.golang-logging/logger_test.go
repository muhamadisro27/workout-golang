package golanglogging

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = logrus.New()

func TestLogger(t *testing.T) {
	logger.Println("Hello World")
}

func TestLevel(t *testing.T) {
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
}

func TestOutput(t *testing.T) {
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Error(err)
	}
	logger.SetOutput(file)

	logger.Info("This is trace")
}

func TestFormatter(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("application.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Error(err)
	}
	logger.SetOutput(file)

	logger.Info("This is trace")
}

type ErrorFields map[string]interface{}

func TestField(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("application.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Error(err)
	}
	defer file.Close()

	logger.SetOutput(file)

	logger.WithField("user", "Isro").Info("Hello World")

	logger.WithFields(logrus.Fields{
		"username": "isro",
		"name":     "Muhamad Isro Sabanur",
	}).Info("Hello World")
}

func TestEntry(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithField("username", "isro")
	entry.Info("Hello World")
}

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(e *logrus.Entry) error {
	fmt.Println("Sample Hook", e.Level, e.Message)

	return nil
}

func TestHook(t *testing.T) {
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.AddHook(&SampleHook{})

	logger.Info("Hello World")
	logger.Warn("Hello World")
	logger.Errorf("Hello World")
}
