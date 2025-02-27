package calls

import (
	"context"
	"log"
	"time"
)

// Calls FetchAll() every hour
func RefreshDB(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Hour):
			log.Println("Refreshing artists")
			err := FetchAll()
			if err != nil {
				log.Printf("Error refreshing artists: %v\n", err)
			}
		case <-ctx.Done():
			log.Println("Stopping refreshDB")
			return
		}
	}
}
