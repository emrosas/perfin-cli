package cmd

import (
	"fmt"
	"strconv"
)

func validateInteger(value string) error {
	if !isInt(value) {
		return fmt.Errorf("please enter a valid integer")
	}
	return nil
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
