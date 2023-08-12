package sakana

import (
	"fmt"
	"reflect"
	"strings"
)

// Usage generates the usage string
func (c *Command) Usage() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var buf strings.Builder

	if len(c.welcome) != 0 {
		fmt.Fprintln(&buf, c.welcome)
		fmt.Fprintln(&buf)
	}

	if !reflect.ValueOf(c.summary).IsZero() {
		fmt.Fprintf(&buf, "usage: %s\n", c.summary.usage)
		fmt.Fprintf(&buf, "    %s\n", c.summary.description)
		fmt.Fprintln(&buf)
	}

	if len(c.options) != 0 {
		fmt.Fprintln(&buf, "options:")
		var maxWidth int
		var existRequired bool
		for _, option := range c.options {
			var width int
			for i, name := range option.names {
				width += len(name)
				if i != 1 {
					width += 2
				}
			}
			if width > maxWidth {
				maxWidth = width
			}
			existRequired = existRequired || option.required
		}
		for _, option := range c.options {
			var width int
			for i, name := range option.names {
				width += len(name)
				if i != 0 {
					fmt.Fprintf(&buf, ", %s", name)
				} else {
					width += 2
					fmt.Fprintf(&buf, "    %s", name)
				}
			}
			fmt.Fprint(&buf, strings.Repeat(" ", maxWidth-width+1))
			if existRequired {
				if option.required {
					fmt.Fprint(&buf, "(required) ")
				} else {
					fmt.Fprint(&buf, "           ")
				}
			}
			fmt.Fprintf(&buf, "- %s\n", option.description)
		}
		fmt.Fprintln(&buf)
	}

	if len(c.subs) != 0 {
		fmt.Fprintln(&buf, "commands:")
		var maxWidth int
		for name := range c.subs {
			if len(name) > maxWidth {
				maxWidth = len(name)
			}
		}
		for name, command := range c.subs {
			fmt.Fprintf(&buf, "    %s%s - %s\n", strings.Repeat(" ", maxWidth-len(name)), name, command.desc)
		}
		fmt.Fprintln(&buf)
	}

	if len(c.examples) != 0 {
		fmt.Fprintln(&buf, "examples:")
		for _, example := range c.examples {
			fmt.Fprintf(&buf, "    %s\n", example.usage)
			fmt.Fprintf(&buf, "        %s\n", example.description)
		}
		fmt.Fprintln(&buf)
	}

	if buf.Len() == 0 {
		return "(no help message)"
	}
	return buf.String()[:buf.Len()-1]
}