package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[1]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("coulnd't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
