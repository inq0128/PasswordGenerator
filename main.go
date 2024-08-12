package main

import (
	"log"
	"math/big"
	"os"
	"strings"

	"crypto/rand"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#;%:?*()@$%^&*-_=+abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#;%:?*()@$%^&*-_=+abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#;%:?*()@$%^&*-_=+"

func randSymbol(n int) string {
	reader := rand.Reader

	b := make([]byte, n)
	for i := range b {
		randRune, _ := rand.Int(reader, big.NewInt(81))
		b[i] = letterBytes[randRune.Int64()]
	}
	return string(b)
}

func makeResponse(chatID telego.ChatID) *telego.SendMessageParams {
	var sb strings.Builder

	for range 4 {
		sb.WriteString("`")
		sb.WriteString(randSymbol(12))
		sb.WriteString("`")
		sb.WriteString("\n")
	}

	response := tu.Message(chatID, sb.String())
	response.ParseMode = "MarkdownV2"

	return response
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	bot, err := telego.NewBot(os.Getenv("TOKEN"), telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatalln(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)

		_, _ = bot.SendMessage(makeResponse(chatID))
	}, th.CommandEqual("start"))

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Start()
}
