package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var std = func() (std *Loggers) {
	std = NewLoggers()
	std.callDepth = 3
	return
}()

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetLevel(l Level) {
	std.SetLevel(l)
}

func Prefix(prefix string) {
	std.Prefix(prefix)
}

func Debug(v ...any) {
	std.Debug(v...)
}

func Info(v ...any) {
	std.Info(v...)
}

func Warning(v ...any) {
	std.Warning(v...)
}

func Error(v ...any) {
	std.Error(v...)
}

func Panic(v any) {
	std.Panic(v)
}

type Loggers struct {
	mu        sync.RWMutex
	level     Level
	loggers   map[Level]*log.Logger
	callDepth int
}

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelPanic
)

func NewLoggers() *Loggers {
	return &Loggers{
		level: LevelInfo,
		loggers: map[Level]*log.Logger{
			LevelInfo:    log.New(os.Stderr, "(info) ", log.Lmsgprefix|log.LstdFlags|log.LUTC),
			LevelWarning: log.New(os.Stderr, "(warning) ", log.Lmsgprefix|log.Lshortfile|log.LstdFlags|log.LUTC),
			LevelError:   log.New(os.Stderr, "(error) ", log.Lmsgprefix|log.Llongfile|log.LstdFlags|log.LUTC),
			LevelPanic:   log.New(os.Stderr, "(panic) ", log.Lmsgprefix|log.Llongfile|log.LstdFlags|log.LUTC),
			LevelDebug:   log.New(os.Stderr, "(debug) ", log.Lmsgprefix|log.Llongfile|log.LstdFlags|log.LUTC),
		},
		callDepth: 2,
	}
}

func (ls *Loggers) SetOutput(output io.Writer) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	for _, l := range ls.loggers {
		l.SetOutput(output)
	}
}

func (ls *Loggers) SetLevel(l Level) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.level = l
}

func (ls *Loggers) Prefix(prefix string) {
	for _, l := range ls.loggers {
		l.SetPrefix(prefix + " " + l.Prefix())
	}
}

func (ls *Loggers) Debug(v ...any) {
	if ls.level <= LevelDebug {
		ls.loggers[LevelDebug].Output(ls.callDepth, fmt.Sprintln(v...))
	}
}

func (ls *Loggers) Info(v ...any) {
	if ls.level <= LevelInfo {
		ls.loggers[LevelInfo].Output(ls.callDepth, fmt.Sprintln(v...))
	}
}

func (ls *Loggers) Warning(v ...any) {
	if ls.level <= LevelWarning {
		ls.loggers[LevelWarning].Output(ls.callDepth, fmt.Sprintln(v...))
	}
}

func (ls *Loggers) Error(v ...any) {
	if ls.level <= LevelError {
		ls.loggers[LevelError].Output(ls.callDepth, fmt.Sprintln(v...))
	}
}

func (ls *Loggers) Panic(v any) {
	if ls.level <= LevelPanic {
		ls.loggers[LevelPanic].Output(ls.callDepth, fmt.Sprintln(v))
	}
	panic(v)
}
