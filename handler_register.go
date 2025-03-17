package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/t57r/gator/internal/database"
)

func hanlderRegister(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("register: expecting username as an argument")
	}

	userName := cmd.args[1]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	fmt.Printf(" * ID:     %v\n", user.ID)
	fmt.Printf(" * Name:   %v\n", user.Name)

	return nil
}
