package consistencytracker 

import (
	"github.com/google/uuid"
	"time"
)
type ActivityBoard struct {
	
	ActivityId uuid.UUID `db:"activity_id"`
	ActivityDate time.Time `db:"activitydate"`
	Completed bool `db:"completed"`
}
