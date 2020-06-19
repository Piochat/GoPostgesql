package models

import (
	"time"
)

//Estudiante Struct
type Estudiante struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdateAt  time.Time
}
