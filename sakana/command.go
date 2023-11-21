package sakana

import (
	"flag"
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
