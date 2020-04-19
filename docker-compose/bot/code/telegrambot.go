package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/shohrukh56/TelegramBot-Go-wiki/googleAPI"
	"log"
	"reflect"
	""
)

func telegramBot() {

	//Create bot
	bot, err := tgbotapi.NewBotAPI("1239310880:AAGC9mOrA3uOeADzFggdZc2MyT9ZPPPXMIM")
	if err != nil {
		panic(err)
	}

	//Set update timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 0

	//Get updates from bot
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//Check if message from user is text
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
		go func() {
			switch update.Message.Text {
			case "/start":

				log.Print("start")
				//Send message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a wikipedia bot, i can search information in a wikipedia, send me something what you want find in Wikipedia.")
				bot.Send(msg)

			case "/number_of_users":

				log.Print("dddd")
			case "/salom":
				msg:=tgbotapi.NewMessage(update.Message.Chat.ID,"hi brother")
				bot.Send(msg)

			default:

				//Set search language
				log.Println("this is == update.Message.Text ",update.Message.Text)
				request := "https://" + "www.google.com/search?q=" + update.Message.Text
				log.Println("1", request)
				//assigning value of answer slice to variable message
				message := wiki.WikipediaAPI(request)
				log.Println("2",message)

				//Loop throug message slice
				for _, val := range message {

					//Send message
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, val)
					log.Println("3",msg)
					bot.Send(msg)
				}
			}
		}()

		} else {

			//Send message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
			bot.Send(msg)
		}
	}
}

func main() {

	apiKey:="AIzaSyDnfEIFVtnfUGbiIMYU1KNUqp3vjgHx0Qg"
	parameters := map[string]string{
		"q": "pizza",
	}
	response, error := serpwow.GetJSON(parameters, apiKey)

	// print the response, or error, if one occurred
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(response)
	}
	for{
		telegramBot()
	}

}
