package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get all users: %w", err)
	}

	for _, user := range users {
		current := ""
		if user.Name == s.config.CurrentUserName {
			current = " (current)"
		}
		fmt.Printf(" * %s%s\n", user.Name, current)
	}

	return nil
}
