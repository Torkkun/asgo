package discord

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	url string = "http://localhost"
)

//デイリー用のデータ
type DailyGatyaResponse struct {
	DayPoint      int    `json:"day_point"`
	PointSum      int    `json:"point_sum"`
	ExecutionDate string `json:"execution_data"`
}

func DailyGatya(s *discordgo.Session) {
	//t := time.NewTicker(time.Minute)
	t := time.NewTicker(time.Second * 3)
	defer t.Stop()
	for range t.C {
		now := time.Now().Format("15:04")
		if now == "07:00" {
			// getrequestを送信
			r := gethttp()
			var daily DailyGatyaResponse
			if err := json.NewDecoder(r).Decode(&daily); err != nil {
				log.Println(err)
				return
			}
			log.Println("rolled_daily")
			//s.ChannelMessageSend("928889658451054684", "test")
			//embededで表示
			return
		}
	}
}

func gethttp() io.Reader {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer res.Body.Close()
	r := io.TeeReader(res.Body, os.Stderr)
	return r
}

func TestTicker(s *discordgo.Session) {
	t := time.NewTicker(time.Duration(5) * time.Second)
	defer t.Stop()
	for range t.C {
		log.Println("test")
		s.ChannelMessageSend("928889658451054684", "test")
	}
}
