package lab

import "time"

type Config struct {
	ID        int
	CreatedAt time.Time
	Desc      string
	Path      string
}
