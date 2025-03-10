package main

import (
	"errors"
	"fmt"
)

func hanlderLogin(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("login: expecting username as an argument")
	}

	userName := cmd.args[1]
	err := s.config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("The user name has been changed to %s\n", userName)
	return nil
}
