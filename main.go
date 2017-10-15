package main

import (
	"flag"

	"os"
	"os/signal"
	"syscall"

	"github.com/HCScorp/bot/handlers"
	"github.com/LorisFriedel/go-bot/bot"
	"github.com/LorisFriedel/go-bot/router"
	log "github.com/sirupsen/logrus"
)

type Arguments struct {
	token string
}

var argToken string

func init() {
	flag.StringVar(&argToken, "t", "", "Discord Authentication Token")
}

func main() {
	// Parse arguments
	flag.Parse()
	argsCli := parseCli()
	argsEnvVar := parsEnvVar()
	args := merge(argsEnvVar, argsCli)

	// Set up
	hcsbot, err := bot.New(args.token)
	if err != nil {
		log.Errorln(err)
		return
	}

	// TODO deport route creation
	ahRoute, err := router.RouteBuilder.Contains("ah").Handler(handlers.NewAh()).Build()
	if err != nil {
		log.Errorln(err)
		return
	}
	hcsbot.Router.AddRoute("ah", ahRoute)

	log.Infoln("The HCS bot is running. Press CTRL-C to exit.")
	waitToBeMurdered()

	// Clean up
	err = hcsbot.Stop()
	if err != nil {
		log.Errorln(err)
		return
	}
}

// Wait for a CTRL-C
func waitToBeMurdered() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

///////////////////////////////////////////
//////////////// Parsing //////////////////
///////////////////////////////////////////

func parsEnvVar() *Arguments {
	return &Arguments{
		token: os.Getenv("TOKEN"),
	}
}

func parseCli() *Arguments {
	return &Arguments{
		token: argToken,
	}
}

// merge aggregate arguments from every given sources. Override are made regarding the method arguments order.
// (the last one can override the first one but not the other way round)
func merge(argsList ...*Arguments) *Arguments {
	result := &Arguments{}
	for _, args := range argsList {
		if args.token != "" {
			result.token = args.token
		}
	}
	return result
}
