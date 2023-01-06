package Bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type (
	CommandRouter struct {
		bot      *Bot
		commands map[string]*Command
	}
)

func (router *CommandRouter) SaveCommand(command string, commandCallback callback, aliases []string) *Command {
	router.commands[command] = NewCommand(command, commandCallback)
	router.commands[command].SetAliases(aliases)
	return router.commands[command]
}

func (router *CommandRouter) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, router.bot.GetPrefix()) {
		return
	}

	commandName, args := ParseCommand(m.Content[1:])
	command := router.commands[commandName]
	ctx := NewContext(s, m, router, command, args)

	if command != nil {
		command.commandCallback(ctx)
		return
	}

	ctx.PrepareMessage(fmt.Sprintf("Invalid command: %s", commandName)).Reply()
}

func (router *CommandRouter) Commands() map[string]*Command {
	return router.commands
}

func (router *CommandRouter) Close() {
	for k := range router.commands {
		delete(router.commands, k)
	}
}

func NewRouter(bot *Bot) (r *CommandRouter) {
	return &(CommandRouter{bot: bot, commands: make(map[string]*Command)})
}
