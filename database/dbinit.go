package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func DbInitialization(db *pgx.Conn) {
	DDLs := []string{ClientsDDL, JuniorDDL}
	for _, ddl := range DDLs {
		_, err := db.Exec(context.Background(), ddl)
		if err != nil {
			log.Fatalf("can't create a table", err)
		}
	}
}
