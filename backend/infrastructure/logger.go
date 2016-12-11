package infrastructure

import (
	"github.com/Sirupsen/logrus"
)

// Logger logs messages.
type Logger struct{}

// Log logs the message. For the time being just the Info method of the logrus
// library is being used. It works like a general logger without any specific
// context like warning or error.
func (logger *Logger) Log(message string) error {
	logrus.Info(message)
	return nil
}
