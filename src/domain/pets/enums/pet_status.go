package pets_enums

import (
	"database/sql/driver"
	"fmt"
)

type PetStatus string

const (
	DRAFT       PetStatus = "DRAFT"
	UNPUBLISHED PetStatus = "UNPUBLISHED"
	PUBLISHED   PetStatus = "PUBLISHED"
)

func (ps *PetStatus) Scan(value interface{}) error {
	*ps = PetStatus(fmt.Sprint(value))
	return nil
}

func (ps PetStatus) Value() (driver.Value, error) {
	return string(ps), nil
}
