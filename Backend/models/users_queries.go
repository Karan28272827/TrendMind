package models

import "time"

type UserQueries struct {
	UserID     string    `json:"user_id"`
	ChatID     string    `json:"chat_id"`
	Query      string    `json:"query"`
	Response   string    `json:"response"`
	TimeStamps time.Time `json:"time_stamps"`
}

type ChatHistory struct {
	UserID     string    `json:"user_id"`
	ChatID     string    `json:"chat_id"`
	ChatName   string    `json:"chat_name"`
	TimeStamps time.Time `json:"time_stamps"`
}

