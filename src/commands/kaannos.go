package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var kaannosInfo = discordgo.ApplicationCommand{
	Name:        "käännös",
	Description: "Käännös mint",
}

func kaannosCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, "https://mxrr.dev/files/vanamehe/käännös.mp4")

	if err != nil {
		log.Printf("Käännös issue: %v", err)
	}
}
