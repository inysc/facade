package facade

type Logger interface {
	Trace(string)
	Tracef(string, ...any)
	Debug(string)
	Debugf(string, ...any)
	Info(string)
	Infof(string, ...any)
	Warn(string)
	Warnf(string, ...any)
	Error(string)
	Errorf(string, ...any)
}

type nopLogger struct{}

func (l nopLogger) Trace(string)          {}
func (l nopLogger) Tracef(string, ...any) {}
func (l nopLogger) Debug(string)          {}
func (l nopLogger) Debugf(string, ...any) {}
func (l nopLogger) Info(string)           {}
func (l nopLogger) Infof(string, ...any)  {}
func (l nopLogger) Warn(string)           {}
func (l nopLogger) Warnf(string, ...any)  {}
func (l nopLogger) Error(string)          {}
func (l nopLogger) Errorf(string, ...any) {}

var l Logger = nopLogger{}

func SetLogger(lg Logger) {
	l = lg
}

func Trace(msg string) {
	l.Trace(msg)
}

func Tracef(msg string, args ...any) {
	l.Tracef(msg, args...)
}

func Debug(msg string) {
	l.Debug(msg)
}

func Debugf(msg string, args ...any) {
	l.Debugf(msg, args...)
}

func Info(msg string) {
	l.Info(msg)
}

func Infof(msg string, args ...any) {
	l.Infof(msg, args...)
}

func Warn(msg string) {
	l.Warn(msg)
}

func Warnf(msg string, args ...any) {
	l.Warnf(msg, args...)
}

func Error(msg string) {
	l.Error(msg)
}

func Errorf(msg string, args ...any) {
	l.Errorf(msg, args...)
}
