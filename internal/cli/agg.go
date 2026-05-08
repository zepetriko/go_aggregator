package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zepetriko/go_aggregator/internal/database"
	"github.com/zepetriko/go_aggregator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("agg <time_between_reqs> arg is required")
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Println("Error scraping feeds: ", err)
		}
	}

	return nil
}

func scrapeFeeds(s *State) error {
	next_feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.Db.MarkFeedFetched(context.Background(), next_feed.ID)
	if err != nil {
		return err
	}

	fetched_feed, err := rss.FetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		return err
	}

	for _, item := range fetched_feed.Channel.Item {
		parsedTime, parseErr := time.Parse(time.RFC1123Z, item.PubDate)

		_, err = s.Db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Url:         item.Link,
				Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
				PublishedAt: sql.NullTime{Time: parsedTime, Valid: parseErr == nil},
				FeedID:      next_feed.ID,
			},
		)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue //duplicate URL, ignore
				}
			}
			log.Printf("error creating post: %v", err)
		}
	}

	return nil
}
