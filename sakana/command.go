package sakana

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"sync"
)

// Command routes handler by command (`FlagSet.Args()[0]`)
type Command struct {
	FlagSet *flag.FlagSet

	mu       sync.RWMutex
	welcome  string
	summary  example
	options  []option
	examples []example

	work Handler
	subs map[string]command
}

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

func (c *Command) Work(work Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.work = work
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

func (c *Command) ServeArgs(w io.Writer, args []string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if w == nil {
		w = output
	}

	if err := c.FlagSet.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Fprint(w, c.Usage())
			return
		}
		UsageError(w, "Failed to parse arguments: "+err.Error())
	}
	args = c.FlagSet.Args()

	if c.work != nil {
		c.work.ServeArgs(w, args)
	}

	c.FlagSet.Args()
	if len(args) > 0 {
		sub, ok := c.subs[args[0]]
		if !ok {
			UsageError(w, fmt.Sprintf("Undefined subcommand (%s).", args[0]))
		}

		sub.h.ServeArgs(w, args[1:])
	}
}
