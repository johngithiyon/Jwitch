package include

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

var Err error

func GetConnection() {
	url := "postgres://john:john@localhost:5432/jwitch"

	Conn,Err = pgx.Connect(context.Background(),url)

	if Err != nil {
		fmt.Println("Database connection is failed")
		return
	} 
	fmt.Println("Database connection  is successful")
}