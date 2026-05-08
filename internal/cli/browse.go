package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/zepetriko/go_aggregator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return errors.New("browse <limit>(optional) takes 1 or none arguments")
	}

	limit := 2
	if len(cmd.Args) > 0 {
		limit, _ = strconv.Atoi(cmd.Args[0])
	}

	posts, err := s.Db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		},
	)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
	}

	return nil
}
