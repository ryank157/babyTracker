package services

import (
	"context"
	"time"

	"github.com/ryank157/babyTracker/backend/db"
)

type EventService interface {
    CreateEvent(ctx context.Context, eventType string, eventTime time.Time, notes *string, mood *db.MoodType) (db.Event, error)
    GetEvent(ctx context.Context, eventID int32) (db.Event, error)
CreateAndAssociateFeedingEvent(ctx context.Context, eventType string, eventTime time.Time, notes *string, mood *db.MoodType, amount float64, feedType string, spitup bool, startTime time.Time, endTime time.Time) (db.FeedingEvent, error)
    // Other methods for updating, deleting, listing, etc.
}
