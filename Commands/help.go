package Commands

import (
	"GoDiscordBot/Bot"
	"github.com/bwmarrin/discordgo"
	"time"
)

func Help(ctx *Bot.Context) {
	ctx.Embed(&discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    ctx.User.Username,
			IconURL: ctx.User.AvatarURL("1024"),
		},
		Color: 0xba000d, // Red
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "I am a field",
				Value: "I am a value",
			},
			{
				Name:  "I am a second field",
				Value: "I am a value",
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	})
}
