package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token string
)

//TOKEN
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	Token = os.Getenv("DGU_TOKEN")
	if Token == "" {
		log.Println("token not found")
		return
	}
}

func CreateNewSession() *discordgo.Session {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil
	}
	return dg
}
