package domain

import "time"

type Bloqueo struct {
	Id      int
	Hora    time.Time
	Usuario Usuario
}
