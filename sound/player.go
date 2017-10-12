package sound

import (
	"github.com/bwmarrin/discordgo"
)

// From https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/main.go
func PlaySound(s *discordgo.Session, guildID, channelID string, sound File) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)

	if err != nil {
		return err
	}

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range sound.Content {
		vc.OpusSend <- buff
	}

	// Stop speaking
	vc.Speaking(false)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return nil
}
