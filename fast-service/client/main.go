package main

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	// TODO
	_ = eg
}
