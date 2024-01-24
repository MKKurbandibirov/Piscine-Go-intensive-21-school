package domain

import "time"

type Anomaly struct {
	SessionID string
	Frequency float64
	Time      time.Time
}
