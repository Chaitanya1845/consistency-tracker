package consistencytracker

import (
	"github.com/google/uuid"
	"time"
)
type Activity struct {
	Id uuid.UUID `db:"id"`
	Name string `db:"name"`
	StartDate time.Time	`db:"startdate"`
	EndDate time.Time `db:"enddate"`
	IsActive bool  `db:"isactive"`
	SuccessDays int `db:"successdays"`
	TotalDays int  `db:"totaldays"`
	ActivityBoard uuid.UUID `db:"activityboard_id"` // struct pending for boardId
	UserId uuid.UUID  `db:"user_id"`
}
