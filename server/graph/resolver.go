package graph

import (
	"sync"
	"time"

	"github.com/TomTomRixRix/Feedbackbuch/server/graph/model"
	"github.com/jinzhu/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Notification struct {
	C         chan *model.Comment
	Timestamp time.Time
}

// Resolver is the Resolver
type Resolver struct {
	DB           *gorm.DB
	ToBeNotified []Notification
	mu           sync.Mutex
}
