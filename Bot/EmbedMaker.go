package Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"reflect"
	"strings"
	"time"
)

type (
	EmbedMaker struct {
		Settings *EmbedSettings
		embed    *discordgo.MessageEmbed
		error    error
	}

	EmbedSettings struct {
		Author    *discordgo.MessageEmbedAuthor
		Color     int
		Image     *discordgo.MessageEmbedImage
		Thumbnail *discordgo.MessageEmbedThumbnail
		Timestamp string
	}
)

func (em *EmbedMaker) Setting(name string, value interface{}) (err error) {
	switch strings.ToLower(name) {
	case "author":
		if reflect.TypeOf(value) == reflect.TypeOf(&discordgo.MessageEmbedAuthor{}) {
			em.Settings.Author = value.(*discordgo.MessageEmbedAuthor)
		} else {
			err = errors.New("value must be a *discordgo.MessageEmbedAuthor")
		}
		break
	case "color":
		if reflect.TypeOf(value).Name() == "int" {
			em.Settings.Color = value.(int)
		} else {
			err = errors.New("value must be an int")
		}
		break
	case "image":
		if reflect.TypeOf(value) == reflect.TypeOf(&discordgo.MessageEmbedImage{}) {
			em.Settings.Image = value.(*discordgo.MessageEmbedImage)
		} else {
			err = errors.New("value must be a *discordgo.MessageEmbedImage")
		}
		break
	case "thumbnail":
		if reflect.TypeOf(value) == reflect.TypeOf(&discordgo.MessageEmbedThumbnail{}) {
			em.Settings.Thumbnail = value.(*discordgo.MessageEmbedThumbnail)
		} else {
			err = errors.New("value must be a *discordgo.MessageEmbedThumbnail")
		}
		break
	case "timestamp":
		if reflect.TypeOf(value).Name() == "string" {
			if value.(string) == "now" {
				em.Settings.Timestamp = time.Now().Format(time.RFC3339)
			} else {
				em.Settings.Timestamp = value.(string)
			}
		} else {
			err = errors.New("value must be a string")
		}
		break
	default:
		err = errors.New("unknown setting " + name)
		break
	}
	return
}

func (em *EmbedMaker) Title(title string) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Title = title
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Description(desc string) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Description = desc
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Author(author *discordgo.MessageEmbedAuthor) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Author = author
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Color(color int) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Color = color
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Image(image *discordgo.MessageEmbedImage) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Image = image
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Thumbnail(thumbnail *discordgo.MessageEmbedThumbnail) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Thumbnail = thumbnail
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Timestamp(timestamp string) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		if timestamp == "now" {
			em.Settings.Timestamp = time.Now().Format(time.RFC3339)
		} else {
			em.Settings.Timestamp = timestamp
		}
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Field(name string, value string) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Fields = append(em.embed.Fields, &discordgo.MessageEmbedField{Name: name, Value: value, Inline: false})
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}
	return em
}

func (em *EmbedMaker) InlineField(name string, value string) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Fields = append(em.embed.Fields, &discordgo.MessageEmbedField{Name: name, Value: value, Inline: true})
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}
	return em
}

func (em *EmbedMaker) Fields(fields ...struct {
	Name  string
	Value string
}) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		for _, field := range fields {
			em.embed.Fields = append(em.embed.Fields, &discordgo.MessageEmbedField{Name: field.Name, Value: field.Value, Inline: false})
		}
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}
	return em
}

func (em *EmbedMaker) InlineFields(fields ...struct {
	Name  string
	Value string
}) *EmbedMaker {
	if em.embed != nil && em.error == nil {
		for _, field := range fields {
			em.embed.Fields = append(em.embed.Fields, &discordgo.MessageEmbedField{Name: field.Name, Value: field.Value, Inline: true})
		}
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}
	return em
}

func (em *EmbedMaker) UseSetting() *EmbedMaker {
	if em.embed != nil && em.error == nil {
		em.embed.Author = em.Settings.Author
		em.embed.Color = em.Settings.Color
		em.embed.Image = em.Settings.Image
		em.embed.Thumbnail = em.Settings.Thumbnail
		em.embed.Timestamp = em.Settings.Timestamp
	} else if em.error != nil {
		em.error = errors.New("missing embed")
	}

	return em
}

func (em *EmbedMaker) Get() *discordgo.MessageEmbed {
	return em.embed
}

func NewEmbedMaker(settings *EmbedSettings) *EmbedMaker {
	return &EmbedMaker{Settings: settings, embed: &discordgo.MessageEmbed{}}
}
