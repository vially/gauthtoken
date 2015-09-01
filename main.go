package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: gauthtoken REFRESH_TOKEN")
		os.Exit(1)
	}

	conf := oauth2.Config{
		ClientID:     "32555940559.apps.googleusercontent.com",
		ClientSecret: "ZmssLNjJy2998hD4CTg2ejr2",
		Endpoint:     google.Endpoint,
		RedirectURL:  "oob",
	}

	refreshToken := &oauth2.Token{
		AccessToken:  "",
		RefreshToken: os.Args[1],
		Expiry:       time.Now().UTC(),
	}

	token, err := conf.TokenSource(oauth2.NoContext, refreshToken).Token()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(token.AccessToken)
}
