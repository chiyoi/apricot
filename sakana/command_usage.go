package sakana

import (
	"fmt"
	"strings"
)

func (c *Command) String() string {
	return fmt.Sprintf("(%s: %s)", c.name, c.summary.description)
}

func (c *Command) UsageString() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var buf strings.Builder
	if len(c.welcome) != 0 {
		fmt.Fprintln(&buf, c.welcome)
		fmt.Fprintln(&buf)
	}

	fmt.Fprintf(&buf, "Usage: %s\n", c.summary.usage)
	fmt.Fprintf(&buf, "    %s\n", c.summary.description)
	fmt.Fprintln(&buf)

	if len(c.options) != 0 {
		fmt.Fprintln(&buf, "Options:")
		var maxWidth int
		for _, option := range c.options {
			var width int
			for i, name := range option.names {
				width += len(name)
				if i > 0 {
					width += 2
				}
			}
			if width > maxWidth {
				maxWidth = width
			}
		}
		for _, option := range c.options {
			var width int
			for i, name := range option.names {
				width += len(name)
				if i != 0 {
					width += 2
					fmt.Fprintf(&buf, ", %s", name)
				} else {
					fmt.Fprintf(&buf, "    %s", name)
				}
			}
			fmt.Fprint(&buf, strings.Repeat(" ", maxWidth-width))
			fmt.Fprintf(&buf, " - %s\n", option.description)
		}
		fmt.Fprintln(&buf)
	}

	if len(c.subs) != 0 {
		fmt.Fprintln(&buf, "Commands:")
		var maxWidth int
		for name := range c.subs {
			if len(name) > maxWidth {
				maxWidth = len(name)
			}
		}
		for name, h := range c.subs {
			padding := strings.Repeat(" ", maxWidth-len(name))
			if cmd, ok := h.(*Command); ok {
				fmt.Fprintf(&buf, "    %s%s - %s\n", padding, name, cmd.summary.description)
			} else {
				fmt.Fprintf(&buf, "    %s%s\n", padding, name)
			}
		}
		fmt.Fprintln(&buf)
	}

	if len(c.examples) != 0 {
		fmt.Fprintln(&buf, "Examples:")
		for _, example := range c.examples {
			fmt.Fprintf(&buf, "    %s\n", example.usage)
			fmt.Fprintf(&buf, "        %s\n", example.description)
		}
		fmt.Fprintln(&buf)
	}

	if buf.Len() == 0 {
		return "(No help message.)"
	}
	return buf.String()[:buf.Len()-1]
}
