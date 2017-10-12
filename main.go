package main

import (
	"os"

	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/thomasmunoz13/bot/cmd"
	"github.com/thomasmunoz13/bot/sound"
)

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateStatus(0, "Ah !")
}

func main() {
	//discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	discord, err := discordgo.New("Bot MzY4MDUwODE2NzgwODYxNDQw.DMEpAA.Tfd-Evoc0zayXxjOoF28CxpDimU")

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	sounds := make(map[string]*sound.File)
	sounds["ah"], err = sound.LoadFile("sounds/ah.mp3", true)

	if err != nil {
		logrus.Fatalf("Could not load sound %s \n", err.Error())
		return
	}

	discord.AddHandler(ready)

	h := cmd.Handler{Sounds: sounds}
	discord.AddHandler(h.Handle)

	// Open the websocket and begin listening.
	err = discord.Open()

	if err != nil {
		logrus.Fatal("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	logrus.Info("Ah Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
