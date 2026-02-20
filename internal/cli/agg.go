package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zepetriko/go_aggregator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("no argument is required for command 'agg'")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
