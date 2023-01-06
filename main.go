package main

import (
	"GoDiscordBot/Bot"
	"GoDiscordBot/Commands"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

// TODO: log system
func main() {
	log := color.New(color.FgBlue, color.Bold)
	log.Print("Reading environment variables... ")
	environment, err := godotenv.Read()
	if err != nil {
		color.Red("✘\nCan't read environment variables, ", err)
		os.Exit(0)
		return
	}
	color.Green("✔")

	if len(environment["PREFIX"]) == 0 {
		environment["PREFIX"] = ">"
	}

	log.Print("Creating discord session... ")
	bot, err := Bot.New(environment["TOKEN"], environment["PREFIX"])
	bot.SetPrefix(environment["PREFIX"])

	if err != nil {
		color.Red("✘\nCan't create discord session, ", err)
		return
	}
	color.Green("✔")

	log.Print("Commands registration... ")
	bot.SaveCommand("help", Commands.Help).SetDesc("Show help")
	color.Green("\u2714")

	log.Print("Adding handler... ")
	bot.Session().AddHandler(bot.MessageCreate)
	bot.Session().AddHandler(ready)
	color.Green("✔")

	log.Print("Opening web socket... ")
	err = bot.Open()

	if err != nil {
		color.Red("✘\nCan't open websocket, ", err)
		return
	}
	color.Green("✔")
	fmt.Println("No error found")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = bot.Close()
	if err != nil {
		color.Red("✘\nA unknown problem occurred while closing bot, ", err)
		return
	}
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	user := r.User

	fmt.Printf("\n\nID: %s\nUsername: %s#%s\nGuilds: %d\n\n\n\n",
		user.ID,
		user.Username,
		user.Discriminator,
		len(r.Guilds))
	fmt.Println(user.Username + " is running. CTRL+C to exit")
}
