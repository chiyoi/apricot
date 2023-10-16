package sakana

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := NewCommand("neko", "neko [nyan|nyan1]", "Nyan~")
	cmd.Welcome("Nyan~")
	cmd.Example("neko nyan", "Nyan~")
	cmd.OptionUsage([]string{"a", "bcd"}, false, "Nyan.")
	cmd.OptionUsage([]string{"e"}, false, "Nyan1.")
	cmd.Command(NewCommand("nyan", "nyan 1 2 3", "Nyan 1 2 3."))
	fmt.Println(cmd)
}
