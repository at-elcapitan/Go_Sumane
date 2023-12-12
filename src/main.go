// Copyright (C) 2023 ElCapitan; pungentee
//
// This file is part of GoSM.
//
// GoSM is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GoSM is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with GoSM.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	dotenv "github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not load .env: %s", err))
	}

	botSettings := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(botSettings)

	if err != nil {
		log.Fatal(fmt.Sprintf("Could not create bot: %s", err))
	}

	bot.Handle("/ping", func(c tele.Context) error {
		return c.Send("Pong!")
	})

	bot.Start()

}
