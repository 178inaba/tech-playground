package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	var count int
	var nextCursor string
	for {
		chs, nc, err := api.GetConversationsContext(
			context.Background(),
			&slack.GetConversationsParameters{
				Limit:  1000,
				Cursor: nextCursor,
				//ExcludeArchived: true,
			},
		)
		if err != nil {
			log.Fatalf("Get conversations: %v.", err)
		}

		count += len(chs)
		for _, ch := range chs {
			fmt.Println(ch.Name)
		}

		if nc == "" {
			break
		}
		nextCursor = nc
	}

	fmt.Println("----------")
	fmt.Printf("Count: %d.\n", count)
}
