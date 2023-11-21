package sakana

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

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
			if _, err := fmt.Fprint(f.Err, c.UsageString()); err != nil {
				ls.Error(err)
				return 2
			}
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
