package cmd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/thomasmunoz13/bot/sound"
)

type Handler struct {
	Sounds map[string]*sound.File
}

var (
	AH string = "ah"
)

func (h Handler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Checks if the message contains "ah"
	if strings.Contains(strings.ToLower(m.Content), AH) {

		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)

		if err != nil {
			// Could not find channel.
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			return
		}

		// Look for the message sender in that guild's current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				err = sound.PlaySound(s, g.ID, vs.ChannelID, *h.Sounds[AH])
				if err != nil {
					fmt.Println("Error playing sound:", err)
				}

				return
			}
		}
	}
}
