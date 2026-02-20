package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("no argument is required for command 'reset'")
	}

	err := s.Db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Table users cleaned!")
	return nil
}
