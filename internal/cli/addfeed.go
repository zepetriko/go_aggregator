package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zepetriko/go_aggregator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed <name> <url> required")
	}

	feedname := cmd.Args[0]
	feedURL := cmd.Args[1]

	currentUser, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.Db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      feedname,
			Url:       feedURL,
			UserID:    currentUser.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Feed created: %+v\n", feed)
	return nil
}
