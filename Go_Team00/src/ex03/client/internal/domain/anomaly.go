package domain

import "time"

type Anomaly struct {
	SessionID string
	Freauency float64
	Time      time.Time
}
