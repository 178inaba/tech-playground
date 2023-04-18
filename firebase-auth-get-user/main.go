package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	ctx := context.Background()

	flag.Parse()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("New Firebase app: %v.", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Get Firebase Authentication client: %v.", err)
	}

	ur, err := client.GetUser(ctx, flag.Arg(0))
	if err != nil {
		log.Fatalf("Get user: %v.", err)
	}

	fmt.Printf("auth.UserRecord: %+v\n", ur)

	fmt.Printf("DisplayName: %s\n", ur.DisplayName)
	fmt.Printf("Email      : %s\n", ur.Email)
	fmt.Printf("PhoneNumber: %s\n", ur.PhoneNumber)
	fmt.Printf("PhotoURL   : %s\n", ur.PhotoURL)
	fmt.Printf("ProviderID : %s\n", ur.ProviderID)
	fmt.Printf("UID        : %s\n", ur.UID)
}
