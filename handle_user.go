package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/remcous/bootdev_gator/internal/database"
)

/*******************************************************************************
*	Login
*******************************************************************************/

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("cannot login, username [%s] not registered", userName)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("The current user has been set to [%s]\n", s.cfg.CurrentUserName)
	return nil
}

/*******************************************************************************
*	Register
*******************************************************************************/

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      userName,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user [%s] was created: %+v\n", userName, s.db)
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
