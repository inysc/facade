package hog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/inysc/facade/infra/util"
)

type event struct {
	*bytes.Buffer
	io.Writer
	lvl  Level
	done func([]byte)
}

var _ Event = (*event)(nil)

func getevent() *event { return eventpl.Get().(*event) }
func freeEvent(e Event) {
	switch e := e.(type) {
	case *event:
		if e.Len() > maxLen {
			e.Reset()
			eventpl.Put(e)
		}
	}
}

func (e *event) Init(ctx context.Context, skip int, lvl uint8, op string, w io.Writer) Event {
	e.lvl = lvl
	e.Writer = w
	appendTime(e.Buffer, time.Now())

	switch lvl {
	case TRACE:
		e.Buffer.WriteString(" TRC ")
	case DEBUG:
		e.Buffer.WriteString(" DBG ")
	case INFO:
		e.Buffer.WriteString(" INF ")
	case WARN:
		e.Buffer.WriteString(" WRN ")
	case ERROR:
		e.Buffer.WriteString(" ERR ")
	case FATAL:
		e.Buffer.WriteString(" FAT ")
	default:
		e.Buffer.WriteByte(' ')
	}

	// 写入调用信息
	appendCaller(e.Buffer, skip)

	// 写入 goid
	// e.WriteString("|goid:")
	// e.Buffer.Write(transNum(runtime.Goid()))
	e.WriteString(" op=")
	e.WriteString(op)
	for _, v := range util.CtxPairs {
		if val := util.CtxString(ctx, v.Key); val != "" {
			e.WriteString("||")
			e.WriteString(v.Val)
			e.WriteString(val)
		}
	}

	return e
}

func (e *event) String(key string, val string) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.WriteString(val)
	return e
}

func (e *event) Strings(key string, val []string) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) != 0 {
		e.WriteString(val[0])
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.WriteString(val[i])
		}
	}
	return e
}

func (e *event) Stringp(key string, val *string) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.WriteString(*val)
	} else {
		e.WriteString("<nil>")
	}
	return e
}

func (e *event) Duration(key string, val time.Duration) Event {
	e.WriteString(key)
	e.WriteString(key)
	e.WriteByte('=')
	e.WriteString(val.String())
	return e
}

func (e *event) Durations(key string, val []time.Duration) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.WriteString(val[0].String())
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.WriteString(val[i].String())
		}
	}
	return e
}

func (e *event) Durationp(key string, val *time.Duration) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.WriteString(val.String())
	} else {
		e.WriteString("<nil>")
	}
	return e
}

func (e *event) Byte(key string, val byte) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.WriteByte(val)
	return e
}

func (e *event) Bytes(key string, val []byte) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(val)

	return e
}

func (e *event) Bytep(key string, val *byte) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.WriteByte(*val)
	} else {
		e.WriteString("<nil>")
	}
	return e
}

func (e *event) Bool(key string, val bool) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.WriteString(strconv.FormatBool(val))
	return e
}

func (e *event) Bools(key string, val []bool) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.WriteString(strconv.FormatBool(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.WriteString(strconv.FormatBool(val[i]))
		}
	}
	return e
}

func (e *event) Boolp(key string, val *bool) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.WriteString(strconv.FormatBool(*val))
	} else {
		e.WriteString("<nil>")
	}
	return e
}

func (e *event) Int(key string, val int) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))
	return e
}

func (e *event) Ints(key string, val []int) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Intp(key string, val *int) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Int8(key string, val int8) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Int8s(key string, val []int8) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Int8p(key string, val *int8) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Int16(key string, val int16) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Int16s(key string, val []int16) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Int16p(key string, val *int16) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Int32(key string, val int32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Int32s(key string, val []int32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Int32p(key string, val *int32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Int64(key string, val int64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Int64s(key string, val []int64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Int64p(key string, val *int64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Uint(key string, val uint) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Uints(key string, val []uint) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Uintp(key string, val *uint) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Uint8(key string, val uint8) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Uint8s(key string, val []uint8) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Uint8p(key string, val *uint8) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Uint16(key string, val uint16) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Uint16s(key string, val []uint16) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Uint16p(key string, val *uint16) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Uint32(key string, val uint32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Uint32s(key string, val []uint32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Uint32p(key string, val *uint32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Uint64(key string, val uint64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(transNum(val))

	return e
}

func (e *event) Uint64s(key string, val []uint64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if len(val) == 0 {
		e.Buffer.Write(transNum(val[0]))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(transNum(val[i]))
		}
	}
	return e
}

func (e *event) Uint64p(key string, val *uint64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.Buffer.Write(transNum(*val))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Any(key string, val any) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	if val != nil {
		e.WriteString(fmt.Sprintf("%+v", val))
	} else {
		e.WriteString("<nil>")
	}
	return e
}

func (e *event) Float32(key string, val float32) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.WriteByte('=')
	e.Buffer.Write(strconv.AppendFloat([]byte{}, float64(val), 'f', -1, 32))

	return e
}

func (e *event) Float32s(key string, val []float32) Event {
	e.WriteString("||")
	e.WriteString(key)
	if len(val) != 0 {
		e.Buffer.Write(strconv.AppendFloat([]byte{'='}, float64(val[0]), 'f', -1, 32))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(strconv.AppendFloat([]byte{}, float64(val[i]), 'f', -1, 32))
		}
	}

	return e
}

func (e *event) Float32p(key string, val *float32) Event {
	e.WriteString("||")
	e.WriteString(key)
	if val != nil {
		e.Buffer.Write(strconv.AppendFloat([]byte{'='}, float64(*val), 'f', -1, 32))
	} else {
		e.WriteString("<nil>")
	}

	return e
}

func (e *event) Float64(key string, val float64) Event {
	e.WriteString("||")
	e.WriteString(key)
	e.Buffer.Write(strconv.AppendFloat([]byte{'='}, val, 'f', -1, 64))
	return e
}

func (e *event) Float64s(key string, val []float64) Event {
	e.WriteString("||")
	e.WriteString(key)
	if len(val) != 0 {
		e.Buffer.Write(strconv.AppendFloat([]byte{'='}, val[0], 'f', -1, 32))
		for i := 1; i < len(val); i++ {
			e.WriteByte(',')
			e.Buffer.Write(strconv.AppendFloat([]byte{'='}, val[i], 'f', -1, 32))
		}
	}

	return e
}

func (e *event) Float64p(key string, val *float64) Event {
	e.WriteString("||")
	e.WriteString(key)
	if val != nil {
		e.Buffer.Write(strconv.AppendFloat([]byte{'='}, *val, 'f', -1, 64))
	} else {
		e.WriteString("=<nil>")
	}

	return e
}

func (e *event) Error(key string, val error) Event {
	e.WriteString("||")
	e.WriteString(key)
	if val != nil {
		e.WriteByte('=')
		e.WriteString(val.Error())
	} else {
		e.WriteString("=<nil>")
	}

	return e
}

func (e *event) IgError(key string, val error) Event {
	if val != nil {
		e.WriteString("||")
		e.WriteString(key)
		e.WriteByte('=')
		e.WriteString(val.Error())
	}

	return e
}

func (e *event) Done()                           { e.write("") }
func (e *event) Msg(msg string)                  { e.write(msg) }
func (e *event) Msgf(format string, args ...any) { e.write(fmt.Sprintf(format, args...)) }

func (e *event) write(msg string) {
	if msg != "" {
		e.WriteString(msg)
	}
	e.WriteByte('\n')
	e.Writer.Write(e.Buffer.Bytes())

	if e.done != nil {
		e.done(e.Buffer.Bytes())
	}

	switch e.lvl {
	case FATAL:
		os.Stderr.Write(e.Buffer.Bytes())
		os.Exit(1)
	case PANIC:
		panic(e.Buffer.String())
	}

	freeEvent(e)
}

type evempty struct{}

func (ne evempty) Init(context.Context, int, uint8, string, io.Writer) Event { return ne }
func (ne evempty) Any(string, any) Event                                     { return ne }
func (ne evempty) Error(string, error) Event                                 { return ne }
func (ne evempty) IgError(string, error) Event                               { return ne }
func (ne evempty) String(string, string) Event                               { return ne }
func (ne evempty) Strings(string, []string) Event                            { return ne }
func (ne evempty) Stringp(string, *string) Event                             { return ne }
func (ne evempty) Duration(string, time.Duration) Event                      { return ne }
func (ne evempty) Durations(string, []time.Duration) Event                   { return ne }
func (ne evempty) Durationp(string, *time.Duration) Event                    { return ne }
func (ne evempty) Byte(string, byte) Event                                   { return ne }
func (ne evempty) Bytes(string, []byte) Event                                { return ne }
func (ne evempty) Bytep(string, *byte) Event                                 { return ne }
func (ne evempty) Bool(string, bool) Event                                   { return ne }
func (ne evempty) Bools(string, []bool) Event                                { return ne }
func (ne evempty) Boolp(string, *bool) Event                                 { return ne }
func (ne evempty) Int(string, int) Event                                     { return ne }
func (ne evempty) Ints(string, []int) Event                                  { return ne }
func (ne evempty) Intp(string, *int) Event                                   { return ne }
func (ne evempty) Int8(string, int8) Event                                   { return ne }
func (ne evempty) Int8s(string, []int8) Event                                { return ne }
func (ne evempty) Int8p(string, *int8) Event                                 { return ne }
func (ne evempty) Int16(string, int16) Event                                 { return ne }
func (ne evempty) Int16s(string, []int16) Event                              { return ne }
func (ne evempty) Int16p(string, *int16) Event                               { return ne }
func (ne evempty) Int32(string, int32) Event                                 { return ne }
func (ne evempty) Int32s(string, []int32) Event                              { return ne }
func (ne evempty) Int32p(string, *int32) Event                               { return ne }
func (ne evempty) Int64(string, int64) Event                                 { return ne }
func (ne evempty) Int64s(string, []int64) Event                              { return ne }
func (ne evempty) Int64p(string, *int64) Event                               { return ne }
func (ne evempty) Uint(string, uint) Event                                   { return ne }
func (ne evempty) Uints(string, []uint) Event                                { return ne }
func (ne evempty) Uintp(string, *uint) Event                                 { return ne }
func (ne evempty) Uint8(string, uint8) Event                                 { return ne }
func (ne evempty) Uint8s(string, []uint8) Event                              { return ne }
func (ne evempty) Uint8p(string, *uint8) Event                               { return ne }
func (ne evempty) Uint16(string, uint16) Event                               { return ne }
func (ne evempty) Uint16s(string, []uint16) Event                            { return ne }
func (ne evempty) Uint16p(string, *uint16) Event                             { return ne }
func (ne evempty) Uint32(string, uint32) Event                               { return ne }
func (ne evempty) Uint32s(string, []uint32) Event                            { return ne }
func (ne evempty) Uint32p(string, *uint32) Event                             { return ne }
func (ne evempty) Uint64(string, uint64) Event                               { return ne }
func (ne evempty) Uint64s(string, []uint64) Event                            { return ne }
func (ne evempty) Uint64p(string, *uint64) Event                             { return ne }
func (ne evempty) Float32(string, float32) Event                             { return ne }
func (ne evempty) Float32s(string, []float32) Event                          { return ne }
func (ne evempty) Float32p(string, *float32) Event                           { return ne }
func (ne evempty) Float64(string, float64) Event                             { return ne }
func (ne evempty) Float64s(string, []float64) Event                          { return ne }
func (ne evempty) Float64p(string, *float64) Event                           { return ne }
func (ne evempty) Done()                                                     {}
func (ne evempty) Msg(string)                                                {}
func (ne evempty) Msgf(string, ...any)                                       {}

// type fpevent struct{}

// func (ne fpevent) Init(bool, int, uint8, io.Writer) Event  { return ne }
// func (ne fpevent) Any(string, any) Event                   { return ne }
// func (ne fpevent) Error(string, error) Event               { return ne }
// func (ne fpevent) IgError(string, error) Event             { return ne }
// func (ne fpevent) String(string, string) Event             { return ne }
// func (ne fpevent) Strings(string, []string) Event          { return ne }
// func (ne fpevent) Stringp(string, *string) Event           { return ne }
// func (ne fpevent) Duration(string, time.Duration) Event    { return ne }
// func (ne fpevent) Durations(string, []time.Duration) Event { return ne }
// func (ne fpevent) Durationp(string, *time.Duration) Event  { return ne }
// func (ne fpevent) Byte(string, byte) Event                 { return ne }
// func (ne fpevent) Bytes(string, []byte) Event              { return ne }
// func (ne fpevent) Bytep(string, *byte) Event               { return ne }
// func (ne fpevent) Bool(string, bool) Event                 { return ne }
// func (ne fpevent) Bools(string, []bool) Event              { return ne }
// func (ne fpevent) Boolp(string, *bool) Event               { return ne }
// func (ne fpevent) Int(string, int) Event                   { return ne }
// func (ne fpevent) Ints(string, []int) Event                { return ne }
// func (ne fpevent) Intp(string, *int) Event                 { return ne }
// func (ne fpevent) Int8(string, int8) Event                 { return ne }
// func (ne fpevent) Int8s(string, []int8) Event              { return ne }
// func (ne fpevent) Int8p(string, *int8) Event               { return ne }
// func (ne fpevent) Int16(string, int16) Event               { return ne }
// func (ne fpevent) Int16s(string, []int16) Event            { return ne }
// func (ne fpevent) Int16p(string, *int16) Event             { return ne }
// func (ne fpevent) Int32(string, int32) Event               { return ne }
// func (ne fpevent) Int32s(string, []int32) Event            { return ne }
// func (ne fpevent) Int32p(string, *int32) Event             { return ne }
// func (ne fpevent) Int64(string, int64) Event               { return ne }
// func (ne fpevent) Int64s(string, []int64) Event            { return ne }
// func (ne fpevent) Int64p(string, *int64) Event             { return ne }
// func (ne fpevent) Uint(string, uint) Event                 { return ne }
// func (ne fpevent) Uints(string, []uint) Event              { return ne }
// func (ne fpevent) Uintp(string, *uint) Event               { return ne }
// func (ne fpevent) Uint8(string, uint8) Event               { return ne }
// func (ne fpevent) Uint8s(string, []uint8) Event            { return ne }
// func (ne fpevent) Uint8p(string, *uint8) Event             { return ne }
// func (ne fpevent) Uint16(string, uint16) Event             { return ne }
// func (ne fpevent) Uint16s(string, []uint16) Event          { return ne }
// func (ne fpevent) Uint16p(string, *uint16) Event           { return ne }
// func (ne fpevent) Uint32(string, uint32) Event             { return ne }
// func (ne fpevent) Uint32s(string, []uint32) Event          { return ne }
// func (ne fpevent) Uint32p(string, *uint32) Event           { return ne }
// func (ne fpevent) Uint64(string, uint64) Event             { return ne }
// func (ne fpevent) Uint64s(string, []uint64) Event          { return ne }
// func (ne fpevent) Uint64p(string, *uint64) Event           { return ne }
// func (ne fpevent) Float32(string, float32) Event           { return ne }
// func (ne fpevent) Float32s(string, []float32) Event        { return ne }
// func (ne fpevent) Float32p(string, *float32) Event         { return ne }
// func (ne fpevent) Float64(string, float64) Event           { return ne }
// func (ne fpevent) Float64s(string, []float64) Event        { return ne }
// func (ne fpevent) Float64p(string, *float64) Event         { return ne }
// func (ne fpevent) Done()                                   { panic("") }
// func (ne fpevent) Msg(string)                              { panic("") }
// func (ne fpevent) Msgf(string, ...any)                     { panic("") }
