package consistencytracker

import (
	"github.com/google/uuid"
)

type User struct {
	
	Id uuid.UUID	`db:"id"`
	Name string		`db:"name"`		 
	Email string	`db:"email"` 
}	
