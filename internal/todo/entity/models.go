// Code generated by sqlc. DO NOT EDIT.

package entity

import (
	"database/sql"
)

type Todo struct {
	ID          int64
	Title       string
	Description sql.NullString
	Status      int32
}
