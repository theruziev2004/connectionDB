package models

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type Users struct {
	ID int64
	Name string
	Surname string
}

func GetUsers(conn *pgx.Conn)  (User []Users, err error) {
	rows, err := conn.Query(context.Background(), `select id, name, surname from clients`)
	if err != nil {
		log.Printf("can't getUsers from clients err is %e", err)
		return User, err
	}
	for rows.Next() {
		Users := Users{}
		err := rows.Scan(
			&Users.ID,
			&Users.Name,
			&Users.Surname)
		if err != nil {
			log.Printf("can't scan %e", err)
		}
		User = append(User, Users)
	}
	if rows.Err() != nil {
		log.Printf("rows err %e", err)
		return nil, rows.Err()
	}
	return User, nil
}
