package cli

import (
	"context"
	"errors"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) > 1 {
		return errors.New("no argument is required for command 'users'")
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) == 0 {
		fmt.Println("No users in the table")
	}
	for _, user := range users {
		if user == s.Config.CurrentUserName {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}
	return nil
}
