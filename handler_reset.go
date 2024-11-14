package main

import (
	"context"
	"fmt"
)

/*******************************************************************************
*	Reset
*******************************************************************************/

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset database: %w", err)
	}

	fmt.Println("successfully reset the database")
	return nil
}
