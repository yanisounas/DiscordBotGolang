package Bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

type (
	Context struct {
		session *discordgo.Session
		message *discordgo.MessageCreate
		router  *CommandRouter

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

func (ctx *Context) Embed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return ctx.session.ChannelMessageSendEmbed(ctx.message.ChannelID, embed)
}

func (ctx *Context) embed() (*discordgo.Message, error) {
	fmt.Println("Embed")
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    ctx.User.Username,
			URL:     "https://google.com",
			IconURL: ctx.User.AvatarURL("1024"),
		},
		Color:       0xba000d, // Red
		Description: "This is a discordgo embed",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "I am a field",
				Value:  "I am a value",
				Inline: true,
			},
			{
				Name:   "I am a second field",
				Value:  "I am a value",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: ctx.User.AvatarURL("2048"),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.User.AvatarURL("2048"),
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "I am an Embed",
	}

	msg, err := ctx.session.ChannelMessageSendEmbed(ctx.message.ChannelID, embed)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return msg, nil
}

func (ctx *Context) Session() *discordgo.Session {
	return ctx.session
}

func (ctx *Context) Message() *discordgo.MessageCreate {
	return ctx.message
}

func (ctx *Context) Router() *CommandRouter {
	return ctx.router
}

func (ctx *Context) GetArgs() []string {
	return ctx.args
}

func (ctx *Context) GetArg(index int) string {
	if len(ctx.args) > index {
		return ctx.args[index]
	}

	return ""
}

func NewContext(session *discordgo.Session, message *discordgo.MessageCreate, router *CommandRouter, command *Command, args []string) (ctx *Context) {
	return &(Context{
		session: session,
		message: message,
		router:  router,
		User:    session.State.User,
		Author:  message.Author,
		command: command,
		args:    args})
}
