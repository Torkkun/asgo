package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//メッセージがbotだった場合無視
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "hi" {
		s.ChannelMessageSend(m.ChannelID, "hello!")
	}
	if m.Content == "whoami" {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ID:%s, NAME:%s", m.Author.ID, m.Author.Username))
	}
	//if m.Content == "mydata" {
	//s.ChannelMessageSendEmbed(m.ChannelID, )
	//}
}
