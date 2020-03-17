package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogrus makes a new Interface backed by a logrus logger
func NewLogrus(level logrus.Level) Logger {
	log := logrus.New()
	log.Out = os.Stderr
	log.Level = level
	return logrusLogger{log}
}

// Logrus wraps an existing Logrus logger.
func Logrus(l *logrus.Logger) Logger {
	return logrusLogger{l}
}

type logrusLogger struct {
	*logrus.Logger
}

func (l logrusLogger) WithField(key string, value interface{}) Logger {
	return logrusEntry{
		Entry: l.Logger.WithField(key, value),
	}
}

func (l logrusLogger) WithFields(fields Fields) Logger {
	return logrusEntry{
		Entry: l.Logger.WithFields(map[string]interface{}(fields)),
	}
}

type logrusEntry struct {
	*logrus.Entry
}

func (l logrusEntry) WithField(key string, value interface{}) Logger {
	return logrusEntry{
		Entry: l.Entry.WithField(key, value),
	}
}

func (l logrusEntry) WithFields(fields Fields) Logger {
	return logrusEntry{
		Entry: l.Entry.WithFields(map[string]interface{}(fields)),
	}
}
