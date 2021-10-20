package main

import "asgo/infra"

var (
	Token string
)

/*
func main() {
	Token := os.Getenv("DGU_TOKEN")
	if Token == "" {
		return
	}
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	//we only care about receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	//Cleanly close down the discord session
	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ignore all message created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping!")
	}
}*/

func main() {
	infra.Letstartsele()

}
