package model

import "time"

// Job is a struct of job
type Job struct {
	JobID       int
	Origin      string
	Destination string
	ShipDate    time.Time
	Budget      int64
	ShipperID   int
}
