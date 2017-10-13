package sound

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

// From https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/main.go
func PlaySound(s *discordgo.Session, guildID, channelID string, file string) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)

	if err != nil {
		return err
	}

	// Start speaking.
	vc.Speaking(true)

	dgvoice.PlayAudioFile(vc, file, make(chan bool))
	// Stop speaking
	vc.Speaking(false)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return nil
}
