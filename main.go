package main

import (
	"GoDiscordBot/Bot"
	"GoDiscordBot/Commands"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	environment, err := godotenv.Read()
	if err != nil {
		return
	}

	bot, err := Bot.New(environment["TOKEN"])
	bot.SetPrefix(environment["PREFIX"])

	if err != nil {
		fmt.Println("can't create discord session:\n", err)
		return
	}

	bot.SaveCommand("ping", Commands.HelloWorld, "pingpong")

	bot.Session().AddHandler(bot.MessageCreate)
	bot.Session().AddHandler(ready)

	err = bot.Open()

	if err != nil {
		fmt.Println("can't open websocket:\n", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	user := r.User

	fmt.Printf("Bot informations:\nID: %s\nUsername: %s#%s\nGuilds: %d\n\n\n\n",
		user.ID,
		user.Username,
		user.Discriminator,
		len(r.Guilds))
	fmt.Println(user.Username + " is running. CTRL+C to exit")
}
