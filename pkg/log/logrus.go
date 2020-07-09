package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

type hook struct{}

func (hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook) Fire(entry *logrus.Entry) error {
	lv := _infoLevel
	var logrusLv string
	var verbose int32
	switch entry.Level {
	case logrus.FatalLevel, logrus.PanicLevel:
		logrusLv = entry.Level.String()
		fallthrough
	case logrus.ErrorLevel:
		lv = _errorLevel
	case logrus.WarnLevel:
		lv = _warnLevel
	case logrus.InfoLevel:
		lv = _infoLevel
	case logrus.DebugLevel:
		// use verbose log replace of debuglevel
		verbose = 10
	}
	args := make([]D, 0, len(entry.Data)+1)
	args = append(args, D{Key: _log, Value: entry.Message})
	for k, v := range entry.Data {
		args = append(args, D{Key: k, Value: v})
	}
	if logrusLv != "" {
		args = append(args, D{Key: "logrus_lv", Value: logrusLv})
	}
	if verbose != 0 {
		V(verbose).Infov(context.Background(), args...)
	} else {
		h.Log(context.Background(), lv, args...)
	}
	return nil
}