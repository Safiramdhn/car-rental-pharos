package helpers

import (
	"log"
	"time"
)

func FormatDate(dateStr string) time.Time {
	const dateLayout = "01/02/2006"

	parsedDate, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		log.Fatalf("Failed to parse date %s: %v", dateStr, err)
	}
	return parsedDate
}
