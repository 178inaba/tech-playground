package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	flag.Parse()
	userID := flag.Arg(0)
	if userID == "" {
		log.Fatal("User ID is required.")
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))
	groups, err := api.GetUserGroupsContext(context.Background(), slack.GetUserGroupsOptionIncludeUsers(true))
	if err != nil {
		log.Fatalf("Get user groups: %v.", err)
	}

	for _, g := range groups {
		for _, u := range g.Users {
			if u == userID {
				fmt.Printf("%s: %s\n", g.Handle, g.Name)
			}
		}
	}
}
