package cmd

import "fmt"

// Check is a struct to help grouping checks before command execution
type Check struct {
	isPositive bool
	message    string
}

// RunChecks runs a list of checks and returns for the first error
func RunChecks(checks []Check) error {
	for _, check := range checks {
		if check.isPositive {
			return fmt.Errorf(check.message)
		}
	}
	return nil
}
