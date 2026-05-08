package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zepetriko/go_aggregator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("follow <url> required")
	}

	feedURL := cmd.Args[0]

	feed, err := s.Db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	feed_follow, err := s.Db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Linked %s with %s\n", feed_follow.FeedName, feed_follow.UserName)
	return nil
}

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("'following' takes no args")
	}

	feeds_user, err := s.Db.GetFeedFollowsForUser(
		context.Background(),
		user.ID,
	)
	if err != nil {
		return err
	}

	for _, feed := range feeds_user {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("unfollow <url> required")
	}

	feedURL := cmd.Args[0]

	feed, err := s.Db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	err = s.Db.DeleteFeedFollowForUser(
		context.Background(),
		database.DeleteFeedFollowForUserParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Unfollowed %s for user %s\n", feed.Name, user.Name)

	return nil
}
