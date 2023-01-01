package Bot

import (
	"github.com/bwmarrin/discordgo"
)

type (
	Context struct {
		session *discordgo.Session
		message *discordgo.MessageCreate

		Author *discordgo.User
		User   *discordgo.User

		command *Command
		args    []string
	}
)

func (ctx *Context) Send(message string) (*discordgo.Message, error) {
	return ctx.session.ChannelMessageSend(ctx.message.ChannelID, message)
}

func (ctx *Context) Reply(message string) (*discordgo.Message, error) {
	return ctx.session.ChannelMessageSendReply(ctx.message.ChannelID,
		message,
		&(discordgo.MessageReference{MessageID: ctx.message.ID, ChannelID: ctx.message.ChannelID, GuildID: ctx.message.GuildID}))
}

func (ctx *Context) GetSession() *discordgo.Session {
	return ctx.session
}

func (ctx *Context) GetMessage() *discordgo.MessageCreate {
	return ctx.message
}

func (ctx *Context) GetArgs() []string {
	return ctx.args
}

func (ctx *Context) GetArg(index int) string {
	//TODO: Check index
	return ctx.args[index]
}

func NewContext(session *discordgo.Session, message *discordgo.MessageCreate, command *Command, args []string) (ctx *Context) {
	return &(Context{
		session: session,
		message: message,
		User:    session.State.User,
		Author:  message.Author,
		command: command,
		args:    args})
}
