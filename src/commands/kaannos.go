package commands

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var kaannosInfo = discordgo.ApplicationCommand{
	Name:        "käännös",
	Description: "Käännös mint",
}

func kaannosCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ChannelID

	log.Println("asd")

	file, err := os.Open("assets/käännös.mp4")
	if err != nil {
		log.Printf("Error reading käännös.mp4 : %v", err)
		return
	}

	v := discordgo.File{
		Name:        "käännös.mp4",
		ContentType: "video/mp4",
		Reader:      file,
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "Se jäättis",
		},
	})

	if err != nil {
		log.Printf("Käännös issue: %v", err)
	}

	s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Files: []*discordgo.File{
			&v,
		},
	})
}
