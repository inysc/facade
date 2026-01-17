package hog

import (
	"context"
	"io"
	"os"
)

type Logger interface {
	Trace(ctx context.Context, op string) Event
	Debug(ctx context.Context, op string) Event
	Info(ctx context.Context, op string) Event
	Warn(ctx context.Context, op string) Event
	Error(ctx context.Context, op string) Event
	Fatal(ctx context.Context, op string) Event
	Panic(ctx context.Context, op string) Event

	Op(ctx context.Context, op string) Event

	SetLevel(Level)
	AddSkip(int)

	NewEvent(context.Context, int, Level, bool, string) Event
}

type nillogger struct{}

var DefaultLogger Logger = nillogger{}

type logger struct {
	io.Writer
	lvl  Level
	skip int
}

func (l *logger) NewEvent(ctx context.Context, skip int, lvl Level, flag bool, op string) Event {
	if lvl >= l.lvl {
		e := getevent()
		e.Init(ctx, l.skip, lvl, op, l.Writer)

		return e
	}
	return evempty{}
}

func (l *logger) Trace(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, TRACE, false, op)
}

func (l *logger) Debug(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, DEBUG, false, op)
}

func (l *logger) Info(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, INFO, false, op)
}

func (l *logger) Warn(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, WARN, false, op)
}

func (l *logger) Error(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, ERROR, false, op)
}

func (l *logger) Fatal(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, FATAL, false, op)
}

func (l *logger) Panic(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, PANIC, false, op)
}

func (l *logger) Op(ctx context.Context, op string) Event {
	return l.NewEvent(ctx, l.skip, OP, true, op)
}

func (l *logger) SetLevel(lvl Level) { l.lvl = lvl }
func (l *logger) AddSkip(skip int)   { l.skip += skip }

func New(lvl Level, ws ...io.Writer) Logger {
	var w io.Writer
	switch len(ws) {
	case 0:
		w = os.Stdout
	case 1:
		w = ws[0]
	default:
		w = io.MultiWriter(ws...)
	}
	return &logger{w, lvl, 5}
}

func Simple(filename string) Logger {
	return New(DEBUG, &LoggerFile{
		Filename:   filename,
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 11,
		LocalTime:  false,
		Compress:   false,
	})
}

func (nillogger) NewEvent(context.Context, int, uint8, bool, string) Event { return evempty{} }
func (nillogger) Trace(ctx context.Context, op string) Event               { return evempty{} }
func (nillogger) Debug(ctx context.Context, op string) Event               { return evempty{} }
func (nillogger) Info(ctx context.Context, op string) Event                { return evempty{} }
func (nillogger) Warn(ctx context.Context, op string) Event                { return evempty{} }
func (nillogger) Error(ctx context.Context, op string) Event               { return evempty{} }
func (nillogger) Fatal(ctx context.Context, op string) Event               { return evempty{} }
func (nillogger) Panic(ctx context.Context, op string) Event               { return evempty{} }
func (nillogger) Op(ctx context.Context, op string) Event                  { return evempty{} }
func (nillogger) SetLevel(Level)                                           {}
func (nillogger) AddSkip(int)                                              {}
