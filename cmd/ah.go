package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/thomasmunoz13/bot/x/mux"
	"github.com/thomasmunoz13/bot/sound"
	"github.com/sirupsen/logrus"
)
var Ah mux.HandlerFunc = func(s *discordgo.Session, msg *discordgo.Message, ctx *mux.Context) {
	// Find the guild for that channel
	g, err := s.State.Guild(ctx.GuildID)

	if err != nil {
		// Could not find guild.
		logrus.Error(err.Error())
		return
	}


	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == msg.Author.ID {

			err = sound.PlaySound(s, g.ID, vs.ChannelID, "sounds/ah.mp3")

			if err != nil {
				logrus.Error(err.Error())
			}

			return
		}
	}
}
