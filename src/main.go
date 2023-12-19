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
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var (
	logger = log.Default()
)

func execScript(db *sql.DB, scriptName string) (sql.Result, error) {
	file, err := os.ReadFile(scriptName)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(string(file))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	_, err = execScript(db, "sqlscripts/init.sql")
	if err != nil {
		logger.Fatal(err)
	}

	bot, err := tele.NewBot(tele.Settings{
		Token: os.Getenv("TOKEN"),
	})
	if err != nil {
		logger.Fatal(err)
	}

	bot.Use(middleware.AutoRespond())

	bot.Handle("/menu", mainMenuHandler)

	bot.Start()
}

// All handlers must be placed here
var (
	getHomeworkButtonHandler = func(c tele.Context) error {
		return c.Send("ДЗ на цей тиждень:")
	}

	newHomeworkButtonHandler = func(c tele.Context) error {
		return c.Send("Надішли мені повідомлення з новим дз")
	}

	checkHomeworkButtonHandler = func(c tele.Context) error {
		return c.Send("Ось ДЗ для перевірки:")
	}

	deleteHomeworkButtonHandler = func(c tele.Context) error {
		return c.Send("Ось дз, яке ти можеш видалити:")
	}

	adminMenuHandler = func(c tele.Context) error {
		adminMenu := &tele.ReplyMarkup{}

		btnCheckHM := adminMenu.Data("Перевірити запити нового ДЗ", "checkHM")
		btnDelHM := adminMenu.Data("Видалити ДЗ з бази данних", "delHM")

		c.Bot().Handle(&btnCheckHM, checkHomeworkButtonHandler)
		c.Bot().Handle(&btnDelHM, deleteHomeworkButtonHandler)

		adminMenu.Inline(adminMenu.Row(btnCheckHM), adminMenu.Row(btnDelHM))

		return c.Send("Меню модератора", adminMenu)
	}

	mainMenuHandler = func(c tele.Context) error {
		// Create new keyboard
		mainMenu := &tele.ReplyMarkup{}

		// Creating buttons for mainMenu keyboard
		btnGetHM := mainMenu.Data("Отримати дз на тиждень", "getHM")
		btnNewHM := mainMenu.Data("Надіслати нове дз", "newHM")
		btnAdminMenu := mainMenu.Data("Меню модератора", "id")

		// Add handlers for buttons to bot (Why we can't set handlers in same place where creating buttons?!?!)
		c.Bot().Handle(&btnGetHM, getHomeworkButtonHandler)
		c.Bot().Handle(&btnNewHM, newHomeworkButtonHandler)
		c.Bot().Handle(&btnAdminMenu, adminMenuHandler)

		// Adding rows to keyboard
		mainMenu.Inline(mainMenu.Row(btnGetHM), mainMenu.Row(btnNewHM), mainMenu.Row(btnAdminMenu))

		return c.Send("Меню бота", mainMenu)
	}
)
