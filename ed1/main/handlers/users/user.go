package users

import (
	"github.com/jackc/pgx/v4"
)

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Lastname string `json:"lastname"`
}

var Conn *pgx.Conn