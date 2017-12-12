package bot

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

		content := strings.Split(m.Content, " ")
		command := content[0]
		fmt.Println(content)

		if command == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Pong")
		}

		if command == "!insult" {
			insult, _ := PersonalAttack(content)
			_, _ = s.ChannelMessageSend(m.ChannelID, insult)
		}

		if command == "!praise" {
			praise, _ := Compliment(content)
			_, _ = s.ChannelMessageSend(m.ChannelID, praise)
		}

		if command == "!joke" {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			c1 := r1.Intn(4)
			fmt.Println(c1)

			if c1 == 0 {
				joke, _ := ChuckJoke()
				_, _ = s.ChannelMessageSend(m.ChannelID, joke)
			}
			if c1 == 1 {
				joke, _ := DadJokes()
				_, _ = s.ChannelMessageSend(m.ChannelID, joke)
			}
			if c1 == 2 {
				joke, _ := MommaJokes()
				_, _ = s.ChannelMessageSend(m.ChannelID, joke)
			}
			if c1 == 3 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Grabbing a dirty joke.. may take a moment")
				joke, _ := DirtyJokes()
				_, _ = s.ChannelMessageSend(m.ChannelID, joke)
			}
		}

		if command == "!angry" {
			rant, _ := Trbmb()
			_, _ = s.ChannelMessageSend(m.ChannelID, rant)
		}

		if command == "!c2f" {
			cc, err := strconv.ParseFloat(content[1], 32)
			if err != nil {
				fmt.Println("C2F Failure")
			}
			c := float32(cc)
			var f float32
			f = ((c * 9) / 5) + 32
			ff := strconv.FormatFloat(float64(f), 'f', 0, 32)
			_, _ = s.ChannelMessageSend(m.ChannelID, content[1]+" Canadian Degrees equals "+ff+" Freedom Degrees")
		}
		if command == "!f2c" {
			ff, err := strconv.ParseFloat(content[1], 32)
			if err != nil {
				fmt.Println("F2C Failure")
			}
			f := float32(ff)
			var c float32
			c = ((f - 32) * 5) / 9
			cc := strconv.FormatFloat(float64(c), 'f', 0, 32)
			_, _ = s.ChannelMessageSend(m.ChannelID, content[1]+" Freedom Degrees equals "+cc+" Canadian Degrees")
		}

	}

}
