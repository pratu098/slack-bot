package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("COmmand Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "Your SLACK_BOT_TOKEN")
	os.Setenv("SLACK_APP_TOKEN", "Your SLACK_APP_TOKEN")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("hello", &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			w.Reply("Hello, how are you? \n How may I help you?")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
