package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// token de acesso
const token string = "XXXXXXXXXXXXX"

// Variavel global para poder ser lida por outros pacotes
var BotID string

func main() {

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botUser, err := discord.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	// Associando o ID do bot a variavel global criada antes
	BotID = botUser.ID

	discord.AddHandler(messageHandler)

	err = discord.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("O bot esta online!")

	// Canal vazio para garantir que o bot estará sempre rodando e esperando algum evento
	<-make(chan struct{})

}
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == BotID {
		return
	}

	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "pong")
	}
}
