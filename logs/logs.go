package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

var std = NewLoggers()

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func Prefix(prefix string) {
	std.Prefix(prefix)
}

func Info(v ...any) {
	std.i.Output(2, fmt.Sprintln(v...))
}

func Warning(v ...any) {
	std.w.Output(2, fmt.Sprintln(v...))
}

func Error(v ...any) {
	std.e.Output(2, fmt.Sprintln(v...))
}

func Debug(v ...any) {
	std.d.Output(2, fmt.Sprintln(v...))
}

func Panic(v any) {
	std.p.Output(2, fmt.Sprintln(v))
	panic(v)
}

type Loggers struct {
	i, w, e, p, f, d *log.Logger
}

func NewLoggers() *Loggers {
	return &Loggers{
		i: log.New(os.Stderr, "(info) ", log.Lmsgprefix|log.LstdFlags|log.LUTC),
		w: log.New(os.Stderr, "(warning) ", log.Lmsgprefix|log.Lshortfile|log.LstdFlags|log.LUTC),
		e: log.New(os.Stderr, "(error) ", log.Lmsgprefix|log.Llongfile|log.LstdFlags|log.LUTC),
		p: log.New(os.Stderr, "(panic) ", log.Lmsgprefix|log.Llongfile|log.LstdFlags|log.LUTC),
		f: log.New(os.Stderr, "(fatal) ", log.Lmsgprefix|log.LstdFlags|log.LUTC),
		d: log.New(os.Stderr, "(debug) ", log.Lmsgprefix|log.Llongfile|log.LstdFlags|log.LUTC),
	}
}

func (ls *Loggers) SetOutput(output io.Writer) {
	for _, l := range []*log.Logger{ls.i, ls.w, ls.e, ls.p, ls.f, ls.d} {
		l.SetOutput(output)
	}
}

func (ls *Loggers) Prefix(prefix string) {
	for _, l := range []*log.Logger{ls.i, ls.w, ls.e, ls.p, ls.f, ls.d} {
		l.SetPrefix(prefix + " " + l.Prefix())
	}
}

func (ls *Loggers) Info(v ...any) {
	ls.i.Output(2, fmt.Sprintln(v...))
}

func (ls *Loggers) Warning(v ...any) {
	ls.w.Output(2, fmt.Sprintln(v...))
}

func (ls *Loggers) Error(v ...any) {
	ls.e.Output(2, fmt.Sprintln(v...))
}

func (ls *Loggers) Debug(v ...any) {
	ls.d.Output(2, fmt.Sprintln(v...))
}

func (ls *Loggers) Panic(v any) {
	ls.p.Output(2, fmt.Sprintln(v))
	panic(v)
}
