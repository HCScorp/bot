package sound

import (
	"encoding/binary"
	"errors"
	"io"
	"os"

	"github.com/jonas747/dca"
	"github.com/sirupsen/logrus"
)

func convert(name string) (*os.File, string, error) {
	logrus.Infof("Loading %s", name)

	opts := dca.StdEncodeOptions
	opts.RawOutput = true
	opts.Bitrate = 60

	// Encoding a file and saving it to disk
	encodeSession, err := dca.EncodeFile(name, opts)

	// Make sure everything is cleaned up, that for example the encoding process if any issues happened isnt lingering around
	defer encodeSession.Cleanup()

	if err != nil {
		return nil, "", err
	}

	tempName := name + "_temp.dca"
	output, err := os.Create(tempName)

	if err != nil {
		return nil, "", err
	}

	io.Copy(output, encodeSession)

	return output, tempName, nil
}

// From https://github.com/bwmarrin/discordgo/blob/master/examples/airhorn/main.go
func LoadFile(name string, toConvert bool) (*File, error) {
	var file *os.File
	var err error

	if toConvert {
		var tempName string
		file, tempName, err = convert(name)
		logrus.Infof("Temp name : %s", tempName)
		//defer os.Remove(tempName)
	} else {
		file, err = os.Open(name)
	}

	buf := make([][]byte, 0)

	if err != nil {
		return nil, err
	}

	var opuslen int16

	for {
		logrus.Info("read binary")
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()

			if err != nil {
				logrus.Error("Error on reading file")
				return nil, err
			}

			return NewFile(name, buf), nil
		}

		if err != nil {
			logrus.Errorf("I don't know %s", err.Error())
			return nil, err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			logrus.Errorf("I really don't know %s", err.Error())
			return nil, err
		}

		// Append encoded pcm data to the buffer.
		buf = append(buf, InBuf)
	}

	return nil, errors.New("Something went really bad in loading file \n")
}
