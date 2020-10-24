package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	authToken := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + authToken)
	if err != nil {
		fmt.Println("Error creating discord section", err)
		return
	}

	dg.AddHandler(messageHandler)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	if err := dg.Open(); err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	fmt.Println("Bot is now running, Press CTRL-C to exit.")
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot's messages
	if s.State.User.ID == m.Author.ID {
		return
	}

	fmt.Println("Message: ", m.Content)

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
