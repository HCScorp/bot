package handlers

import (
	"sync"

	"github.com/HCScorp/bot/sound"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Ah struct {
	mu *sync.Mutex
	// TODO ah sound cached
}

func NewAh() *Ah {
	return &Ah{
		mu: &sync.Mutex{},
	}
}

func (a *Ah) Handle(ds *discordgo.Session, mc *discordgo.MessageCreate, ch *discordgo.Channel) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Find the guild for that channel
	g, err := ds.State.Guild(ch.GuildID)

	if err != nil {
		// Could not find guild.
		logrus.Error(err.Error())
		return
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == mc.Message.Author.ID {
			err = sound.PlaySound(ds, g.ID, vs.ChannelID, "sounds/ah.mp3")

			if err != nil {
				logrus.Error(err.Error())
			}

			return
		}
	}
}
