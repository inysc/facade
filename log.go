package facade

import "context"

type (
	Level = uint8

	Logger interface {
		Trace(string)
		Tracef(string, ...any)
		TraceT(context.Context, string, ...any)
		Debug(string)
		Debugf(string, ...any)
		DebugT(context.Context, string, ...any)
		Info(string)
		Infof(string, ...any)
		InfoT(context.Context, string, ...any)
		Warn(string)
		Warnf(string, ...any)
		WarnT(context.Context, string, ...any)
		Error(string)
		Errorf(string, ...any)
		ErrorT(context.Context, string, ...any)
		Fatal(string)
		Fatalf(string, ...any)
		FatalT(context.Context, string, ...any)

		SetLevel(Level)
		SetTrace(string)
	}
)

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type nopLogger struct{}

func (l nopLogger) Trace(string)                           {}
func (l nopLogger) Tracef(string, ...any)                  {}
func (l nopLogger) TraceT(context.Context, string, ...any) {}
func (l nopLogger) Debug(string)                           {}
func (l nopLogger) Debugf(string, ...any)                  {}
func (l nopLogger) DebugT(context.Context, string, ...any) {}
func (l nopLogger) Info(string)                            {}
func (l nopLogger) Infof(string, ...any)                   {}
func (l nopLogger) InfoT(context.Context, string, ...any)  {}
func (l nopLogger) Warn(string)                            {}
func (l nopLogger) Warnf(string, ...any)                   {}
func (l nopLogger) WarnT(context.Context, string, ...any)  {}
func (l nopLogger) Error(string)                           {}
func (l nopLogger) Errorf(string, ...any)                  {}
func (l nopLogger) ErrorT(context.Context, string, ...any) {}
func (l nopLogger) Fatal(string)                           {}
func (l nopLogger) Fatalf(string, ...any)                  {}
func (l nopLogger) FatalT(context.Context, string, ...any) {}
func (l nopLogger) SetLevel(Level)                         {}
func (l nopLogger) SetTrace(string)                        {}

var l Logger = nopLogger{}

func SetLogger(lg Logger) {
	l = lg
}

func Trace(msg string)                                    { l.Trace(msg) }
func Tracef(msg string, args ...any)                      { l.Tracef(msg, args...) }
func TraceT(ctx context.Context, msg string, args ...any) { l.TraceT(ctx, msg, args...) }

func Debug(msg string)                                    { l.Debug(msg) }
func Debugf(msg string, args ...any)                      { l.Debugf(msg, args...) }
func DebugT(ctx context.Context, msg string, args ...any) { l.DebugT(ctx, msg, args...) }

func Info(msg string)                                    { l.Info(msg) }
func Infof(msg string, args ...any)                      { l.Infof(msg, args...) }
func InfoT(ctx context.Context, msg string, args ...any) { l.InfoT(ctx, msg, args...) }

func Warn(msg string)                                    { l.Warn(msg) }
func Warnf(msg string, args ...any)                      { l.Warnf(msg, args...) }
func WarnT(ctx context.Context, msg string, args ...any) { l.WarnT(ctx, msg, args...) }

func Error(msg string)                                    { l.Error(msg) }
func Errorf(msg string, args ...any)                      { l.Errorf(msg, args...) }
func ErrorT(ctx context.Context, msg string, args ...any) { l.ErrorT(ctx, msg, args...) }

func Fatal(msg string)                                    { l.Error(msg) }
func Fatalf(msg string, args ...any)                      { l.Errorf(msg, args...) }
func FatalT(ctx context.Context, msg string, args ...any) { l.ErrorT(ctx, msg, args...) }
