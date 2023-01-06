package Commands

import (
	"GoDiscordBot/Bot"
)

func Help(ctx *Bot.Context) {
	ctx.PrepareEmbed().Title("Test").Description("Test").Field("FieldTest", "Value").Reply()
}
