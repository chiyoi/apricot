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
	name    string
	summary example
	FlagSet *flag.FlagSet

	mu       sync.RWMutex
	welcome  string
	options  []option
	examples []example

	work []Handler
	subs map[string]*Command
}

var _ Handler = (*Command)(nil)

type example struct {
	usage       string
	description string
}

type option struct {
	names       []string
	required    bool
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
func (c *Command) Command(cmd *Command) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.subs == nil {
		c.subs = map[string]*Command{}
	}

	c.subs[cmd.FlagSet.Name()] = cmd
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
		UsageError(f.Err, "Failed to parse arguments: "+err.Error(), c.UsageString())
		return 1
	}
	args = c.FlagSet.Args()

	for _, h := range c.work {
		ret := h.ServeArgs(f, args)
		if ret != Continue {
			return ret
		}
	}

	if len(args) > 0 {
		sub, ok := c.subs[args[0]]
		if !ok {
			UsageError(f.Err, fmt.Sprintf("Undefined subcommand (%s).", args[0]), c.UsageString())
			return 1
		}
		return sub.ServeArgs(f, args[1:])
	}
	return 0
}
