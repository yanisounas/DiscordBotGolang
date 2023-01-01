package Bot

import (
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

	//TODO: Check if command exists

	ctx := NewContext(s, m, command, args)
	command.commandCallback(ctx)
}

func (router *CommandRouter) Close() {
	for k := range router.commands {
		delete(router.commands, k)
	}
}

func NewRouter(bot *Bot) (r *CommandRouter) {
	return &(CommandRouter{bot: bot, commands: make(map[string]*Command)})
}
