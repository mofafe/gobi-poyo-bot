package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN is null")
	}

	d, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	d.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Username == os.Getenv("POYO_CLIENT_ID") {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			if len(m.Attachments) > 0 {
				s.ChannelMessageSend(m.ChannelID, m.Content+"ぽよ"+"\nファイルの送信はできないぽよ")
			} else {
				s.ChannelMessageSend(m.ChannelID, m.Content+"ぽよ")
			}
		}
	})

	err = d.Open()
	if err != nil {
		log.Fatal(err)
	}
	d.Close()

	log.Println("ボットが起動しました。Ctrl+Cで終了します。")
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot
}
