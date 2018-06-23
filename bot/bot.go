package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mndrix/rand"
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

		if command == "!help" {
			output := `I know many commands here are some of them!
					!ping   (will respond pong) 
					!online (will respond with how many ppl and who is online on play.crafttheory.net)
					!insult @name (I will insult the person named) 
					!praise @name (I will praise the person named) 
					!joke (I'll tell you a funny joke:  Chuck Norris Jokes, Momma Jokes, Dad Jokes, Dirty Jokes)
					!thatreally (I'll tell you what really blanks my blank)  
					!c2f ## (I'll convert Canadian Temps to Freedom Temps)  
                                        !f2c ## (I'll convert Freedom Temps to Canadian Temps)  
                                        !flip heads/tails (I'll flip a coin and tell if your right 
					!cat (will return a random kitty picture)
					!dog <breed> (will return a random dog pic, you can also include breed to narrow it down)

			`
			_, _ = s.ChannelMessageSend(m.ChannelID, output)
		}

		if command == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Pong")
		}
		if command == "!online" {
			playerlist, playercount, _ := WhoIsOnline()
			playerscountstring := "**Number of Active Players:** " + strconv.Itoa(playercount)
			playersliststring := "**Playing on CT Main:** " + strings.Join(playerlist, " **,** ")
			_, _ = s.ChannelMessageSend(m.ChannelID, playerscountstring)
			_, _ = s.ChannelMessageSend(m.ChannelID, playersliststring)
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
			c1 := rand.Intn(4)
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

		if command == "!thatreally" {
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
		if command == "!cat" {
			url, _ := CatPics()
			_, _ = s.ChannelMessageSend(m.ChannelID, url)
		}
		if command == "!dog" {
			if len(content) > 1 {
				url, _ := DogPics(content[1])
				_, _ = s.ChannelMessageSend(m.ChannelID, url)
			} else {
				url, _ := DogPics("empty")
				_, _ = s.ChannelMessageSend(m.ChannelID, url)
			}
		}
		if command == "!flip" {
			coin := rand.Intn(2)
			guess := content[1]
			answer := "blank"
			if coin > 0 {
				answer = "heads"
			} else {
				answer = "tails"
			}
			if guess != "heads" && guess != "tails" {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Sorry my coin only has heads & tails on it.. not: "+content[1])
			} else {
				if guess != "heads" && guess != "tails" {
					fmt.Println("err")
				} else {
					if answer != guess {
						_, _ = s.ChannelMessageSend(m.ChannelID, "Sorry wrong guess, it was "+answer+" :poop:  :poop: ")
					} else {
						_, _ = s.ChannelMessageSend(m.ChannelID, "Congrats!  It was "+answer+"   :cookie: :cookie: ")
					}
				}

			}
		}
	}
}
