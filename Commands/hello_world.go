package Commands

import "GoDiscordBot/Bot"

func HelloWorld(ctx *Bot.Context) {
	ctx.Reply("Hello, " + ctx.GetArg(0) + "!")
}
