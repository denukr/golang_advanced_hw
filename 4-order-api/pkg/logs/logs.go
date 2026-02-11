package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Используем псевдоним log для logrus для удобства
var log = logrus.New()

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	l := &Logger{logrus.New()}
	l.Init()
	return l
}

func (l *Logger) Init() {
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)
}
