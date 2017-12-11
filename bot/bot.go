package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/monkeyworknet/monkeybot/config"
)

var BotID string
var GoBot *discordgo.Session

func Start() {
	fmt.Println("Starting Bot..")

	GoBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error at Token Stage")
		return
	}

	u, err := GoBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error Getting User Info")
	}

	BotID = u.ID
	GoBot.AddHandler(messageHandler)
	err = GoBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error When Connecting")
		return
	}

	fmt.Println("Bot Is Running!")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {

		if m.Author.ID == BotID {
			return
		}

		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Pong")
		}

		if m.Content == "!joke" {
			joke, _ := Insults()
			_, _ = s.ChannelMessageSend(m.ChannelID, joke)

		}
	}

}

func Stop() {
	fmt.Println("Stopping Bot..")
	GoBot.Close()
}
