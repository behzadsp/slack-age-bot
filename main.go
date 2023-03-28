package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2167027336758-5006356306487-ta3wL0ztyfqJjX24OqfbKXfc")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A050P9AF16Y-5017902077061-4645f6356e704fed8bfa30eab3ba24346f99af7816031db19bb5e4c3fd95958e")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Examples:    []string{"my yob is 2023"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error!")
			}
			age := time.Now().Year() - yob
			res := fmt.Sprintf("Your age is: %d", age)
			response.Reply(res)
		},
	})

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
