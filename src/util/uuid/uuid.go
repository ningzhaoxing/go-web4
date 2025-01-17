package uuid

import (
	uuid "github.com/satori/go.uuid"
)

func GetUuid() (string, error) {
	u := uuid.NewV4()
	return u.String(), nil
}
