package Bot

import (
	"github.com/bwmarrin/discordgo"
)

type (
	Bot struct {
		session *discordgo.Session
		prefix  string "!"
		router  *CommandRouter

		User     *discordgo.User
		Username string
		ID       string
	}
)

func (bot *Bot) Session() *discordgo.Session {
	return bot.session
}

func (bot *Bot) Open() (err error) {
	err = bot.session.Open()
	if err != nil {
		return
	}

	bot.User = bot.session.State.User
	bot.Username = bot.User.Username
	bot.ID = bot.User.ID
	return
}

func (bot *Bot) Close() (err error) {
	bot.router.Close()
	return bot.Session().Close()
}

func (bot *Bot) SaveCommand(command string, commandCallback callback, aliases ...string) *Command {
	return bot.router.SaveCommand(command, commandCallback, aliases)
}

func (bot *Bot) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	bot.router.MessageCreate(s, m)
}

func (bot *Bot) SetPrefix(prefix string) {
	bot.prefix = prefix
}

func (bot *Bot) GetPrefix() string {
	return bot.prefix
}

func New(token string) (bot *Bot, err error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot = &(Bot{session: session})
	bot.router = NewRouter(bot)

	return
}
