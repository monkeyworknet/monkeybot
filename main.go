package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/monkeyworknet/monkeybot/bot"
	"github.com/monkeyworknet/monkeybot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error Loading Config Functions")
		return
	}
	bot.Start()

	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Exit Normally.
	bot.Stop()
}
