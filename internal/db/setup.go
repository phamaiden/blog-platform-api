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
	//defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error pinging db: %s", err)
	}

	return conn, nil
}
