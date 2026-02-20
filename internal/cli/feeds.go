package cli

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("'feeds' takes no arguments")
	}

	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return fmt.Errorf("No feeds found")
	}

	for _, f := range feeds {
		fmt.Printf("Name: %s\n", f.Name)
		fmt.Printf("URL: %s\n", f.Url)
		fmt.Printf("Created By: %s\n", f.UserName)
		fmt.Println("--------")
	}

	return nil
}
