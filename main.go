package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

func main() {
	settings := tb.Settings{
		Token:       os.Getenv("BOT_TOKEN"),
		Synchronous: true,
		Verbose:     true,
	}
	tgBot, err := tb.NewBot(settings)
	if err != nil {
		fmt.Println(err)
		panic("can't create bot")
	}
	tgBot.Handle(tb.OnText, func(m *tb.Message) {
		message := m.Text
		tgBot.Send(m.Sender, message)
	})
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var u tb.Update
		if err = json.Unmarshal([]byte(req.Body), &u); err == nil {
			tgBot.ProcessUpdate(u)
		}
		return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil

	})
}
