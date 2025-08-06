package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Init(dbUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %s", err)
	}
	defer conn.Close(context.Background())

	return conn, nil
}
