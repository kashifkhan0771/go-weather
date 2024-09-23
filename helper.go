package go_weather

import (
	"time"
)

// Define the minimum allowed date (2010-01-01)
var minAllowedDate = time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)

// validateQuery ensure the required query in the option is not empty
func validateQuery(q string) bool {
	return q != ""
}

// Ensure the date is on or after January 1st, 2010
func validateDate(date time.Time) bool {
	if date.IsZero() {
		return false
	}

	// Check if the input date is on or after 2010-01-01
	return !date.Before(minAllowedDate)
}
