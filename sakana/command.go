package sakana

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sync"
)

const (
	Continue = -1
)

// Command routes handler by command (`FlagSet.Args()[0]`)
type Command struct {
	FlagSet *flag.FlagSet

	mu       sync.RWMutex
	welcome  string
	summary  example
	options  []option
	examples []example

	work []Handler
	subs map[string]command
}

var _ Handler = (*Command)(nil)

type command struct {
	h    Handler
	desc string
}

type example struct {
	usage       string
	description string
}

type option struct {
	names       []string
	required    bool
	description string
}

func NewCommand(name string) *Command {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.Usage = func() {}
	return &Command{FlagSet: fs}
}

// Welcome registers a welcome message at the beginning of Usage
func (c *Command) Welcome(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.welcome = msg
}

func (c *Command) Summary(usage string, description string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.summary = example{usage, description}
}

// OptionUsage registers a option to display in Usage
func (c *Command) OptionUsage(names []string, required bool, description string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.options = append(c.options, option{names, required, description})
}

// Example registers an example
func (c *Command) Example(usage string, description string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.examples = append(c.examples, example{usage, description})
}

func (c *Command) Work(h Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.work = append(c.work, h)
}

// Command registers or updates a subcommand
func (c *Command) Command(name string, description string, h Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subs == nil {
		c.subs = map[string]command{}
	}

	c.subs[name] = command{h, description}
}

func (c *Command) ServeArgs(w ResponseWriter, args []string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if w == nil {
		w = NewResponseWriter(os.Stdout, os.Stderr)
	}

	if err := c.FlagSet.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Fprint(w.Err(), c.Usage())
			return 0
		}
		UsageError(w, "Failed to parse arguments: "+err.Error())
		fmt.Fprintln(w.Err(), c.Usage())
		return 1
	}
	args = c.FlagSet.Args()

	for _, h := range c.work {
		ret := h.ServeArgs(w, args)
		if ret != Continue {
			return ret
		}
	}

	if len(args) > 0 {
		sub, ok := c.subs[args[0]]
		if !ok {
			UsageError(w, fmt.Sprintf("Undefined subcommand (%s).", args[0]))
			fmt.Fprintln(w.Err(), c.Usage())
			return 1
		}

		return sub.h.ServeArgs(w, args[1:])
	}
	return 0
}
