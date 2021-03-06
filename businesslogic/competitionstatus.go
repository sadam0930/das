package businesslogic

import "time"

const (
	COMPETITION_STATUS_PRE_REGISTRATION    = 1
	COMPETITION_STATUS_OPEN_REGISTRATION   = 2
	COMPETITION_STATUS_CLOSED_REGISTRATION = 3
	COMPETITION_STATUS_IN_PROGRESS         = 4
	COMPETITION_STATUS_PROCESSING          = 5
	COMPETITION_STATUS_CLOSED              = 6
	COMPETITION_STATUS_CANCELLED           = 7
)

type CompetitionStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type ICompetitionStatusRepository interface {
	GetCompetitionStatus() ([]CompetitionStatus, error)
}
