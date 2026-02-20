package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("username is required")
	}

	username := cmd.Args[0]
	_, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	if err := s.Config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("Current user set to %s\n", username)
	return nil
}
