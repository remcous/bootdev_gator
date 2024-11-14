package main

import (
	"fmt"
)

/*******************************************************************************
*	Commands
*******************************************************************************/

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if f, ok := c.registeredCommands[cmd.Name]; ok {
		return f(s, cmd)
	}

	return fmt.Errorf("command [%s] is not supported", cmd.Name)
}

/*******************************************************************************
*	Command
*******************************************************************************/

type command struct {
	Name string
	Args []string
}