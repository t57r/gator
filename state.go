package main

import (
	"errors"
	"fmt"

	"github.com/t57r/gator/internal/config"
	"github.com/t57r/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	hanlders map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.hanlders[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, exist := c.hanlders[cmd.name]
	if !exist {
		return errors.New(fmt.Sprintf("Can't find handler for %s", cmd.name))
	}
	return handler(s, cmd)
}
