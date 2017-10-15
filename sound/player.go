package sound

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

// From https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/main.go
func PlaySound(vc *discordgo.VoiceConnection, file string) (err error) {
	vc.Speaking(true) // Start speaking
	dgvoice.PlayAudioFile(vc, file, make(chan bool))
	vc.Speaking(false) // Stop speaking
	return nil
}
