package discord

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func DailyGatya(s *discordgo.Session) {
	t := time.NewTicker(duration())
	defer t.Stop()
	for range t.C {
		log.Println("rolled_daily")
		//embededで表示
	}
}

//毎朝7時
func duration() time.Duration {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 7, 0, 0, 0, t.Location())
	if t.After(n) {
		n = n.Add(24 * time.Hour)
	}
	duration := n.Sub(t)
	return duration
}

func TestTicker(s *discordgo.Session) {
	t := time.NewTicker(time.Duration(5) * time.Second)
	defer t.Stop()
	for range t.C {
		log.Println("test")
		s.ChannelMessageSend("928889658451054684", "test")
	}
}
