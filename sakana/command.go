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

var (
	_ Handler = (*Command)(nil)
)

// Command routes handler by command (`FlagSet.Args()[0]`)
type Command struct {
	name    string
	summary example
	FlagSet *flag.FlagSet

	mu       sync.RWMutex
	welcome  string
	options  []option
	examples []example

	work []Handler
	subs map[string]Handler
}

type example struct {
	usage       string
	description string
}

type option struct {
	names       []string
	description string
}

func NewCommand(name string, usage string, description string) *Command {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.Usage = func() {}
	return &Command{
		name:    name,
		summary: example{usage, description},
		FlagSet: fs,
	}
}

// Welcome registers a welcome message at the beginning of Usage
func (c *Command) Welcome(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.welcome = msg
}

// OptionUsage registers a option to display in Usage
func (c *Command) OptionUsage(names []string, description string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.options = append(c.options, option{names, description})
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
func (c *Command) Command(name string, h Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subs == nil {
		c.subs = map[string]Handler{}
	}

	c.subs[name] = h
}

func (c *Command) ServeArgs(f Files, args []string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if f.In == nil {
		f.In = os.Stdin
	}
	if f.Out == nil {
		f.Out = os.Stdout
	}
	if f.Err == nil {
		f.Err = os.Stderr
	}

	if err := c.FlagSet.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Fprint(f.Err, c.UsageString())
			return 0
		}
		UsageError(f.Err, err.Error(), c.UsageString())
		return 1
	}
	args = c.FlagSet.Args()

	for _, h := range c.work {
		ret := h.ServeArgs(f, args)
		if ret != Continue {
			return ret
		}
	}

	if len(args) == 0 {
		UsageError(f.Err, "Subcommand is needed.", c.UsageString())
		return 1
	}

	sub, ok := c.subs[args[0]]
	if !ok {
		UsageError(f.Err, fmt.Sprint("Undefined subcommand. ", "args[0]: ", args[0]), c.UsageString())
		return 1
	}
	return sub.ServeArgs(f, args[1:])
}
