package Bot

import (
	"strings"
)

type (
	callback func(ctx *Context)

	Command struct {
		command         string
		commandCallback callback
		description     string
		aliases         []string
		nbArgument      uint
	}
)

func ParseCommand(command string) (string, []string) {
	c := strings.Split(command, " ")
	return c[0], c[1:]
}

func (c *Command) SetDesc(description string) *Command {
	c.description = description
	return c
}

func (c *Command) GetDesc() string {
	return c.description
}

func (c *Command) SetAlias(aliase string) *Command {
	c.aliases = append(c.aliases, aliase)
	return c
}

func (c *Command) SetAliases(aliases []string) *Command {
	for _, alias := range aliases {
		c.SetAlias(alias)
	}
	return c
}

func (c *Command) GetAliases() []string {
	return c.aliases
}

func NewCommand(command string, commandCallback callback) *Command {
	return &(Command{command: command, commandCallback: commandCallback})
}
