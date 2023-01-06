package Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"reflect"
)

type (
	Context struct {
		Session    *discordgo.Session
		Message    *discordgo.MessageCreate
		router     *CommandRouter
		Author     *discordgo.User
		User       *discordgo.User
		EmbedMaker *EmbedMaker
		command    *Command
		args       []string
		message    interface{}
	}
)

func (ctx *Context) Send() (*discordgo.Message, error) {
	if reflect.TypeOf(ctx.message) == reflect.TypeOf(&discordgo.MessageEmbed{}) {
		return ctx.RawEmbed(ctx.EmbedMaker.UseSetting().Get())
	} else if reflect.TypeOf(ctx.message) == reflect.TypeOf("") {
		return ctx.SendMessage(ctx.message.(string))
	}

	return nil, errors.New("unknown error")
}

func (ctx *Context) Reply() (*discordgo.Message, error) {
	ref := &discordgo.MessageReference{MessageID: ctx.Message.ID, ChannelID: ctx.Message.ChannelID, GuildID: ctx.Message.GuildID}
	if reflect.TypeOf(ctx.message) == reflect.TypeOf(&discordgo.MessageEmbed{}) {
		return ctx.RawEmbedReply(ctx.EmbedMaker.UseSetting().Get(), ref)
	} else if reflect.TypeOf(ctx.message) == reflect.TypeOf("") {
		return ctx.SendReply(ctx.message.(string), ref)
	}

	return nil, errors.New("unknown error")
}

func (ctx *Context) SendMessage(message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, message)
}

func (ctx *Context) SendReply(message string, reference *discordgo.MessageReference) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendReply(ctx.Message.ChannelID, message, reference)
}

func (ctx *Context) RawEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}

func (ctx *Context) RawEmbedReply(embed *discordgo.MessageEmbed, reference *discordgo.MessageReference) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendEmbedReply(ctx.Message.ChannelID, embed, reference)
}

func (ctx *Context) PrepareMessage(message string) *Context {
	ctx.message = message
	return ctx
}

func (ctx *Context) PrepareEmbed() *Context {
	ctx.EmbedMaker.Setting("author", &discordgo.MessageEmbedAuthor{Name: ctx.User.Username, IconURL: ctx.User.AvatarURL("1024")})
	ctx.EmbedMaker.Setting("color", 0xba000d)
	ctx.EmbedMaker.Setting("timestamp", "now")
	ctx.message = &discordgo.MessageEmbed{}
	return ctx
}

func (ctx *Context) Title(title string) *Context {
	ctx.EmbedMaker.Title(title)
	return ctx
}

func (ctx *Context) Description(desc string) *Context {
	ctx.EmbedMaker.Description(desc)
	return ctx
}

func (ctx *Context) Image(image *discordgo.MessageEmbedImage) *Context {
	ctx.EmbedMaker.Image(image)
	return ctx
}

func (ctx *Context) Thumbnail(thumbnail *discordgo.MessageEmbedThumbnail) *Context {
	ctx.EmbedMaker.Thumbnail(thumbnail)
	return ctx
}

func (ctx *Context) Field(name string, value string) *Context {
	ctx.EmbedMaker.Field(name, value)
	return ctx
}

func (ctx *Context) InlineField(name string, value string) *Context {
	ctx.EmbedMaker.InlineField(name, value)
	return ctx
}

func (ctx *Context) Fields(fields ...struct {
	Name  string
	Value string
}) *Context {
	ctx.EmbedMaker.Fields(fields...)
	return ctx
}

func (ctx *Context) InlineFields(fields ...struct {
	Name  string
	Value string
}) *Context {
	ctx.EmbedMaker.InlineFields(fields...)
	return ctx
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
		Session:    session,
		Message:    message,
		router:     router,
		User:       session.State.User,
		Author:     message.Author,
		EmbedMaker: NewEmbedMaker(&EmbedSettings{}),
		command:    command,
		args:       args})
}
