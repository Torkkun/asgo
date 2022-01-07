package discord

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func DailyGatya(s *discordgo.Session) {
	//t := time.NewTicker(time.Minute)
	t := time.NewTicker(time.Second * 3)
	defer t.Stop()
	for range t.C {
		now := time.Now().Format("15:04")
		if now == "07:00" {
			log.Println("rolled_daily")
			//s.ChannelMessageSend("928889658451054684", "test")
			//embededで表示
		}
	}
}

func TestTicker(s *discordgo.Session) {
	t := time.NewTicker(time.Duration(5) * time.Second)
	defer t.Stop()
	for range t.C {
		log.Println("test")
		s.ChannelMessageSend("928889658451054684", "test")
	}
}
