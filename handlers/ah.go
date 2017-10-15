package handlers

import (
	"sync"

	"github.com/HCScorp/bot/sound"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type Ah struct {
	mu        *sync.Mutex
	vc        *discordgo.VoiceConnection
	channelId string
	guildID   string
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
		log.Error(err)
		return // Could not find guild.
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == mc.Message.Author.ID {
			// TODO move connection/disconnection logic to Bot type

			if !(a.vc != nil && a.vc.Ready && a.guildID == g.ID && a.channelId == vs.ChannelID) {
				if a.vc != nil {
					a.vc.Disconnect()
				}

				// Join the provided voice channel.
				a.vc, err = ds.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)
				if err != nil {
					log.Error(err)
					return
				}

				a.guildID = g.ID
				a.channelId = vs.ChannelID
			}

			err = sound.PlaySound(a.vc, "sounds/ah.mp3")
			if err != nil {
				log.Error(err)
				return
			}

			// TODO Timeout
			// Disconnect from the provided voice channel.
			// a.vc.Disconnect()

			return
		}
	}
}
