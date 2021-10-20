package infra

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

//TOKEN
func init() {
	Token := os.Getenv("DGU_TOKEN")
	if Token == "" {
		return
	}
}
func CreateNewSession() *discordgo.Session {
	Token := os.Getenv("DGU_TOKEN")
	if Token == "" {
		fmt.Println("token not found")
		return nil
	}
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil
	}
	return dg
}
