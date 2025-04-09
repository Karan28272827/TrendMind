package utils

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

func InitOAuth() {
	//to initiate OAuth for a particular provided (google here)
	goth.UseProviders(
		google.New(
			os.Getenv("OAUTH_CLIENT_ID"),
			os.Getenv("OAUTH_CLIENT_SECRET"),
			"http://localhost:8080/auth/google/callback", //url for the callback after auth on google's side
			"email", "profile", //details we want from google about the user
		),
		github.New(
			os.Getenv("GITHUB_CLIENT_ID"),
			os.Getenv("GITHUB_CLIENT_SECRET"),
			"http://localhost:8080/auth/github/callback",
			"user:email",
		),
	)
}
