package main

import (
	"context"
	"fmt"
)

/*******************************************************************************
*	Handler
*******************************************************************************/

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to retrieve feed, %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}
