package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zepetriko/go_aggregator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("only username argument is required")
	}

	username := cmd.Args[0]

	user, err := s.Db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      username,
		},
	)
	if err != nil {
		return err
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("Current user set to %s\n", username)
	return nil
}
