package facade

import "context"

type (
	Level = uint8

	Logger interface {
		Trace(string)
		TraceT(context.Context, string, ...any)
		Debug(string)
		DebugT(context.Context, string, ...any)
		Info(string)
		InfoT(context.Context, string, ...any)
		Warn(string)
		WarnT(context.Context, string, ...any)
		Error(string)
		ErrorT(context.Context, string, ...any)
		Fatal(string)
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
func (l nopLogger) TraceT(context.Context, string, ...any) {}
func (l nopLogger) Debug(string)                           {}
func (l nopLogger) DebugT(context.Context, string, ...any) {}
func (l nopLogger) Info(string)                            {}
func (l nopLogger) InfoT(context.Context, string, ...any)  {}
func (l nopLogger) Warn(string)                            {}
func (l nopLogger) WarnT(context.Context, string, ...any)  {}
func (l nopLogger) Error(string)                           {}
func (l nopLogger) ErrorT(context.Context, string, ...any) {}
func (l nopLogger) Fatal(string)                           {}
func (l nopLogger) FatalT(context.Context, string, ...any) {}
func (l nopLogger) SetLevel(Level)                         {}
func (l nopLogger) SetTrace(string)                        {}

var l Logger = nopLogger{}

func SetLogger(lg Logger) {
	l = lg
}

func Trace(msg string)                                    { l.Trace(msg) }
func TraceT(ctx context.Context, msg string, args ...any) { l.TraceT(ctx, msg, args...) }

func Debug(msg string)                                    { l.Debug(msg) }
func DebugT(ctx context.Context, msg string, args ...any) { l.DebugT(ctx, msg, args...) }

func Info(msg string)                                    { l.Info(msg) }
func InfoT(ctx context.Context, msg string, args ...any) { l.InfoT(ctx, msg, args...) }

func Warn(msg string)                                    { l.Warn(msg) }
func WarnT(ctx context.Context, msg string, args ...any) { l.WarnT(ctx, msg, args...) }

func Error(msg string)                                    { l.Error(msg) }
func ErrorT(ctx context.Context, msg string, args ...any) { l.ErrorT(ctx, msg, args...) }

func Fatal(msg string)                                    { l.Error(msg) }
func FatalT(ctx context.Context, msg string, args ...any) { l.ErrorT(ctx, msg, args...) }
